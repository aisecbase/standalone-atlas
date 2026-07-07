---
atlas_id: AML.T0048.004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут эксфильтровать ИИ-артефакты, чтобы украсть интеллектуальную собственность и причинить экономический ущерб организации-жертве. Проприетарные обучающие данные дорого собирать и размечать, поэтому...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: AI Intellectual Property Theft
subtechnique_count: 0
subtechnique_of: AML.T0048
tactics:
    - AML.TA0011
title: Кража интеллектуальной собственности ИИ
url: /techniques/AML.T0048.004/
---

Злоумышленники могут эксфильтровать ИИ-артефакты, чтобы украсть интеллектуальную собственность и причинить экономический ущерб организации-жертве.

Проприетарные обучающие данные дорого собирать и размечать, поэтому они могут быть целью для [эксфильтрации](/tactics/AML.TA0010) и кражи.

Провайдеры AIaaS взимают плату за использование своих API.

Злоумышленник, укравший модель через [эксфильтрацию](/tactics/AML.TA0010) или с помощью [извлечения ИИ-модели](/techniques/AML.T0024.002), получает неограниченную возможность использовать этот сервис без оплаты владельцу интеллектуальной собственности.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0048/"><span class="relation-id">AML.T0048</span><strong>Внешний ущерб</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0048.000/"><span class="relation-id">AML.T0048.000</span><strong>Финансовый ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.001/"><span class="relation-id">AML.T0048.001</span><strong>Репутационный ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.002/"><span class="relation-id">AML.T0048.002</span><strong>Общественный вред</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.003/"><span class="relation-id">AML.T0048.003</span><strong>Ущерб пользователям</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать кражу интеллектуальной собственности.</p></a>
<a class="relation-item" href="/mitigations/AML.M0012/"><span class="relation-id">AML.M0012</span><strong>Шифрование чувствительной информации</strong><p>Защищайте артефакты машинного обучения с помощью шифрования.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>Отказ от развертывания моделей на периферийных устройствах снижает потенциальный доступ злоумышленника к моделям или ИИ-артефактам.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0011 Воздействие</span><p>Реплицировав модель с высокой точностью, исследователи показали, что злоумышленник может украсть модель и нарушить права организации-жертвы на интеллектуальную собственность.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0011 Воздействие</span><p>Эксфильтрированные данные могут включать чувствительные или частные данные, например артефакты ML-моделей, хранящиеся в Google Drive.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0011 Воздействие</span><p>Получив полный доступ к моделям, злоумышленник мог украсть ценную интеллектуальную собственность в виде ИИ-моделей.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0011 Воздействие</span><p>Получив полный доступ к моделям, злоумышленник получает ценную интеллектуальную собственность организации.</p></a>
<a class="relation-item" href="/studies/AML.CS0056/"><span class="relation-id">AML.CS0056</span><strong>Кампании по дистилляции моделей, нацеленные на Anthropic Claude</strong><span class="relation-meta">Актор: DeepSeek, Moonshot AI, MiniMax / Тактика: AML.TA0011 Воздействие</span><p>DeepSeek, Moonshot AI и MiniMax переняли возможности Claude с помощью дистилляции, затратив лишь малую часть стоимости разработки собственных моделей. Их интересовали наиболее отличительные возможности Claude, включая агентное рассуждение, использование инструментов и генерацию кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0011 Воздействие</span><p>Восстановленные модели представляли собой проприетарные ИИ-артефакты Google Photos. Злоумышленник или конкурент мог использовать извлеченные модели для изучения, повторного использования или воспроизведения возможностей Google Photos, снижая затраты на самостоятельную разработку аналогичных функций.</p></a>
</div>
