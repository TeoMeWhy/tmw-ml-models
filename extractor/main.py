# %%
import os
import sqlalchemy
import pandas as pd
import numpy as np
import time

import dotenv

dotenv.load_dotenv(dotenv.find_dotenv())

ACCESS_TOKEN    = os.getenv("DATABRICKS_TOKEN")
SERVER_HOSTNAME = os.getenv("DATABRICKS_SERVER_HOSTNAME")
HTTP_PATH       = os.getenv("DATABRICKS_HTTP_PATH")
CATALOG         = os.getenv("DATABRICKS_CATALOG")
SCHEMA          = os.getenv("DATABRICKS_SCHEMA")

def main():
    url = f"databricks://token:{ACCESS_TOKEN}@{SERVER_HOSTNAME}?" + f"http_path={HTTP_PATH}&catalog={CATALOG}&schema={SCHEMA}"
    engine_dbricks = sqlalchemy.create_engine(url=url, echo=True)

    query = """
    SELECT *

    FROM models

    WHERE descLabel = 1
    AND descModelName = 'feature_store.upsell.churn'

    QUALIFY row_number() OVER (PARTITION BY idCliente ORDER BY dtRef DESC) = 1
    """

    with engine_dbricks.connect() as conn:
        df = pd.read_sql_query(query, conn)

    df['nrProbNorm'] = (df['nrProbLabel'] - df['nrProbLabel'].min()) / (df['nrProbLabel'].max() - df['nrProbLabel'].min())

    df = df.sort_values(by="nrProbLabel").reset_index(drop=True)
    df['nrProbaRank'] = df.index + 1
    df['nrProbaRank'] = df['nrProbaRank'] / df.shape[0]

    engine_sqlite = sqlalchemy.create_engine("sqlite:///../data/database.db")
    df.to_sql("churn_score", engine_sqlite, index=False, if_exists='replace')


if __name__ == "__main__":
    while True:
        main()
        time.sleep(60*60*24)