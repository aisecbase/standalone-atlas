---
atlas_id: AML.T0035
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут собирать ИИ-артефакты для эксфильтрации или использования при подготовке атак на ИИ. ИИ-артефакты включают модели и наборы данных, а также другие телеметрические данные, возникающие при...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 4
source_name: AI Artifact Collection
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Сбор ИИ-артефактов
url: /techniques/AML.T0035/
---

Злоумышленники могут собирать ИИ-артефакты для [эксфильтрации](/tactics/AML.TA0010) или использования при [подготовке атак на ИИ](/tactics/AML.TA0001). ИИ-артефакты включают модели и наборы данных, а также другие телеметрические данные, возникающие при взаимодействии с моделью.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничение публикации артефактов может снизить способность злоумышленника собирать артефакты модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать или ограничивать сбор ИИ-артефактов в системе жертвы.</p></a>
<a class="relation-item" href="/mitigations/AML.M0012/"><span class="relation-id">AML.M0012</span><strong>Шифрование чувствительной информации</strong><p>Защищайте артефакты машинного обучения с помощью шифрования.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>Отказ от развертывания моделей на периферийных устройствах уменьшает поверхность атаки и может предотвратить сбор артефактов злоумышленником.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0009 Сбор материалов</span><p>Команда нашла файл целевой ML-модели и необходимые обучающие данные.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0009 Сбор материалов</span><p>Злоумышленник может искать в системе жертвы частные и проприетарные данные, включая артефакты ML-моделей. Jupyter Notebook [позволяют выполнять shell-команды](https://colab.research.google.com/github/jakevdp/PythonDataScienceHandbook/blob/master/notebooks/01.05-IPython-And-Shell-Commands.ipynb). В этом примере смонтированный Drive проверяется на наличие checkpoint-файлов моделей PyTorch: &gt; /content/drive/MyDrive/models/checkpoint.pt</p></a>
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0009 Сбор материалов</span><p>Злоумышленники могут собирать ИИ-артефакты, включая продакшен-модели и данные. Исследователи наблюдали рабочие продакшен-нагрузки нескольких организаций из разных отраслей.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0009 Сбор материалов</span><p>Исследователи собрали артефакты моделей TensorFlow Lite из нескольких мест в APK, включая незашифрованные ресурсы приложения, файлы, встроенные в нативную библиотеку, и каталоги, специфичные для приложения.</p></a>
</div>
