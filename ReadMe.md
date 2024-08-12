# Téo Me Why Models

Um serviço para apoiar o projeto do Téo me Why, integrando modelos de ML criados no Databricks com nosso escossistema nas lives.

## Extrator

Script Python responsável para extrair os dados escorados do Databricks e armazená-los localmente.

Nessa etapa também realizamos a padronização do score usando um rankeamento percentil, isto é, nõa usamos apenas o score puro do modelo, mas sim tranformação para deixá-lo entre 0 e 1.

## API

Uma pequena API criada em GoLang com Gin para fornecer acesso aos dados escorados para demais serviços por meio de um endpoint.