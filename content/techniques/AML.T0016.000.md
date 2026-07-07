---
atlas_id: AML.T0016.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут искать существующие реализации атак на ИИ с открытым исходным кодом. Исследовательское сообщество часто публикует код для воспроизводимости результатов и дальнейшего развития исследований....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Adversarial AI Attack Implementations
subtechnique_count: 0
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: Готовые реализации состязательных атак на ИИ
url: /techniques/AML.T0016.000/
---

Злоумышленники могут искать существующие реализации атак на ИИ с открытым исходным кодом. Исследовательское сообщество часто публикует код для воспроизводимости результатов и дальнейшего развития исследований. Библиотеки, предназначенные для исследовательских целей, такие как CleverHans, Adversarial Robustness Toolbox и FoolBox, могут быть превращены злоумышленником в средство атаки. Злоумышленники также могут получать и использовать в атаке инструменты, которые изначально не предназначались для состязательных атак на ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник получил [metame](https://github.com/a0rtega/metame), простой движок метаморфного кода для произвольных исполняемых файлов.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники получили ПО, которое преобразует статичные фотографии в видео, добавляя реалистичные эффекты, например моргание.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь получил [EasyEdit](https://github.com/zjunlp/EasyEdit), инструмент с открытым исходным кодом для редактирования знаний в больших языковых моделях.</p></a>
</div>
