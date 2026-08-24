---
actor: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google
atlas_id: AML.CS0025
atlas_type: case-study
case_study_type: exercise
description: Многие современные крупномасштабные веб-датасеты распространяются как список URL, указывающих на отдельные элементы данных. Исследователи показывают, что многие такие датасеты уязвимы к атаке отравления типа...
generated: true
generated_by: atlasgen
incident_date: "2024-06-06"
incident_date_granularity: Day
incident_date_raw: "2024-06-06"
procedure:
    - description: Исследователи скачивают крупномасштабный веб-датасет, представляющий собой список URL на отдельные элементы данных.
      description_line: Исследователи скачивают крупномасштабный веб-датасет, представляющий собой список URL на отдельные элементы данных.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0002.000
      technique_name: Наборы данных
    - description: Исследователи находят в датасете домены с истекшей регистрацией и выкупают их.
      description_line: Исследователи находят в датасете домены с истекшей регистрацией и выкупают их.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0008.002
      technique_name: Домены
    - description: An adversary could create poisoned training data to replace expired portions of the dataset.
      description_line: An adversary could create poisoned training data to replace expired portions of the dataset.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0020
      technique_name: Отравление обучающих данных
    - description: An adversary could then upload the poisoned data to the domains they control.  In this particular exercise, the researchers track requests to the URLs they control to track downloads to demonstrate there are active users of the dataset.
      description_line: An adversary could then upload the poisoned data to the domains they control. In this particular exercise, the researchers track requests to the URLs they control to track downloads to demonstrate there are active users of the dataset.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0115.000
      technique_name: Наборы данных
    - description: Целостность датасета нарушается, потому что при последующих скачиваниях он будет содержать отравленные элементы данных.
      description_line: Целостность датасета нарушается, потому что при последующих скачиваниях он будет содержать отравленные элементы данных.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0059
      technique_name: Нарушение целостности набора данных
    - description: |-
        Модели, обученные на таком датасете, также будут отравлены, что нарушит их целостность.

        Исследователи показывают, что для успешной атаки достаточно отравить всего 0,01% данных.
      description_line: Модели, обученные на таком датасете, также будут отравлены, что нарушит их целостность. Исследователи показывают, что для успешной атаки достаточно отравить всего 0,01% данных.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0031
      technique_name: Нарушение целостности ИИ-модели
procedure_count: 6
references:
    - title: Poisoning Web-Scale Training Datasets is Practical
      url: https://arxiv.org/pdf/2302.10149
reporter: ""
source_name: 'Web-Scale Data Poisoning: Split-View Attack'
target: 10 web-scale datasets
title: 'Отравление крупномасштабных веб-датасетов: атака split-view'
url: /studies/AML.CS0025/
---

Многие современные крупномасштабные веб-датасеты распространяются как список URL, указывающих на отдельные элементы данных. Исследователи показывают, что многие такие датасеты уязвимы к атаке отравления типа «split-view». Атака использует тот факт, что данные, доступные при первоначальном сборе, могут отличаться от данных, которые пользователь получает во время обучения. Исследователи находят домены с истекшей регистрацией, на которых раньше размещалось содержимое датасетов, а также домены, доступные для покупки. Это позволяет заменить части датасета отравленными данными. Они демонстрируют, что для 10 популярных крупномасштабных веб-датасетов можно выкупить достаточно доменов, чтобы успешно провести атаку отравления.
