---
atlas_id: AML.T0044
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут получить полный доступ к ИИ-модели в режиме «белого ящика». Это означает, что злоумышленник полностью знает архитектуру модели, её параметры и онтологию классов. Он может эксфильтрировать модель,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 4
source_name: Full AI Model Access
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0000
title: Полный доступ к ИИ-модели
url: /techniques/AML.T0044/
---

Злоумышленники могут получить полный доступ к ИИ-модели в режиме «белого ящика». Это означает, что злоумышленник полностью знает архитектуру модели, её параметры и онтологию классов. Он может эксфильтрировать модель, чтобы [создавать состязательные данные](/techniques/AML.T0043) и [проверять атаки](/techniques/AML.T0042) в офлайн-среде, где его поведение сложно обнаружить.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0000/"><span class="relation-id">AML.TA0000</span><strong>Доступ к ИИ-модели</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа к моделям и данным в состоянии покоя может помочь предотвратить полный доступ к модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>Отказ от распространения модели в ПО для периферийных устройств может ограничить способность злоумышленника получить полный доступ к модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Это дало исследователям полный доступ к ML-модели, хотя и в скомпилированном бинарном виде.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Сотрудники начали пользоваться этой организацией на Hugging Face и загрузили в неё приватные модели. Как владелец учётной записи Hugging Face, исследователь имел полный доступ ко всем этим загруженным моделям с правами на чтение и запись.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Это дало исследователям полный доступ к моделям. Были выявлены модели для разных сценариев применения, включая:</p><ul><li>распознавание удостоверений личности;</li><li>распознавание лиц;</li><li>распознавание объектов;</li><li>различные задачи обработки естественного языка.</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Эксфильтрация файлов моделей из APK дала исследователям полный доступ к ИИ-моделям Google Photos, включая модели для таких задач, как обнаружение лиц, обнаружение объектов, сегментация, оценка глубины, оценка качества изображений и обнаружение размытия.</p></a>
</div>
