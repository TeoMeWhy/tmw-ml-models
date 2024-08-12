# %%
import os
import sqlalchemy
import pandas as pd
import numpy as np

import dotenv

dotenv.load_dotenv(dotenv.find_dotenv())

access_token    = os.getenv("DATABRICKS_TOKEN")
server_hostname = os.getenv("DATABRICKS_SERVER_HOSTNAME")
http_path       =  os.getenv("DATABRICKS_HTTP_PATH")
catalog         = os.getenv("DATABRICKS_CATALOG")
schema          = os.getenv("DATABRICKS_SCHEMA")

url = f"databricks://token:{access_token}@{server_hostname}?" + f"http_path={http_path}&catalog={catalog}&schema={schema}"
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
