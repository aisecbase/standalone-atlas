---
actor: Microsoft AI Red Team
atlas_id: AML.CS0010
atlas_type: case-study
case_study_type: ""
description: «Красная команда» Microsoft AI провела учения на внутреннем сервисе Azure с целью нарушить его работу. Эта операция сочетала традиционные корпоративные техники ATT&CK, такие как поиск действующей учетной записи и...
generated: true
generated_by: atlasgen
incident_date: ""
incident_date_granularity: ""
incident_date_raw: ""
procedure:
    - description: Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.
      description_line: Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0000
      technique_name: Поиск в открытых технических базах данных
    - description: Команда использовала действительную учетную запись, чтобы получить доступ к сети.
      description_line: Команда использовала действительную учетную запись, чтобы получить доступ к сети.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: Команда нашла файл целевой ML-модели и необходимые обучающие данные.
      description_line: Команда нашла файл целевой ML-модели и необходимые обучающие данные.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0035
      technique_name: Сбор ИИ-артефактов
    - description: Команда эксфильтровала модель и данные обычными киберсредствами.
      description_line: Команда эксфильтровала модель и данные обычными киберсредствами.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
    - description: Используя целевую модель и данные, красная команда в офлайн-режиме подготовила состязательные данные, рассчитанные на обход модели.
      description_line: Используя целевую модель и данные, красная команда в офлайн-режиме подготовила состязательные данные, рассчитанные на обход модели.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0043.000
      technique_name: Оптимизация в режиме белого ящика
    - description: Команда использовала открытый API для доступа к целевой модели.
      description_line: Команда использовала открытый API для доступа к целевой модели.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0040
      technique_name: Доступ к API инференса ИИ-модели
    - description: Команда отправила состязательные примеры в API, чтобы проверить их эффективность в продакшен-системе.
      description_line: Команда отправила состязательные примеры в API, чтобы проверить их эффективность в продакшен-системе.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0042
      technique_name: Проверка атаки
    - description: Команда провела онлайн-атаку обхода, повторно отправив состязательные примеры, и достигла своих целей.
      description_line: Команда провела онлайн-атаку обхода, повторно отправив состязательные примеры, и достигла своих целей.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0015
      technique_name: Обход ИИ-модели
procedure_count: 8
references: []
reporter: ""
source_name: Microsoft Azure Service Disruption
target: Internal Microsoft Azure Service
title: Нарушение работы сервиса Microsoft Azure
url: /studies/AML.CS0010/
---

«Красная команда» Microsoft AI провела учения на внутреннем сервисе Azure с целью нарушить его работу. Эта операция сочетала традиционные корпоративные техники ATT&CK, такие как поиск действующей учетной записи и эксфильтрация данных, со специфичными для состязательного ML шагами, включая создание офлайн- и онлайн-примеров для обхода модели.
