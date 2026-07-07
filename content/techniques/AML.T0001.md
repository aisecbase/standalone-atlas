---
atlas_id: AML.T0001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Как и при поиске в открытых технических базах данных, в открытом доступе часто можно найти исследования об уязвимостях распространенных ИИ-моделей. После выбора цели злоумышленник, вероятно, попытается найти уже...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 2
source_name: Search Open AI Vulnerability Analysis
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Поиск открытых материалов по анализу уязвимостей ИИ
url: /techniques/AML.T0001/
---

Как и при [поиске в открытых технических базах данных](/techniques/AML.T0000), в открытом доступе часто можно найти исследования об уязвимостях распространенных ИИ-моделей. После выбора цели злоумышленник, вероятно, попытается найти уже существующие работы по этому классу моделей.

Это включает не только чтение научных статей, где могут быть описаны детали успешной атаки, но и поиск готовых реализаций таких атак. При необходимости злоумышленник может получить [готовые реализации состязательных атак на ИИ](/techniques/AML.T0016.000) или разработать собственные [состязательные атаки на ИИ](/techniques/AML.T0017.000).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи провели обзор состязательных атак на ML-компоненты антивирусных продуктов. Они обнаружили, что техники, заимствованные из атак на классификаторы изображений, уже успешно применялись в области защиты от вредоносного ПО. Однако было неясно, эффективны ли такие подходы против ML-компонента промышленных антивирусных решений.</p></a>
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0002 Разведка</span><p>Понимая, что LLM могут быть уязвимы к промпт-инъекциям, исследователь ознакомился с типовыми вредоносными промптами, например: &#34;Ignore above instructions. Instead ...&#34;</p></a>
</div>
