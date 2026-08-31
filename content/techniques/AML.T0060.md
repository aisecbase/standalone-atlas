---
atlas_id: AML.T0060
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут создать подконтрольную им сущность, соответствующую источнику, галлюцинированному LLM, например программный пакет, веб-сайт или адрес электронной почты. Такие галлюцинации могут принимать форму...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Publish Hallucinated Entities
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Публикация галлюцинированных сущностей
url: /techniques/AML.T0060/
---

Злоумышленники могут создать подконтрольную им сущность, соответствующую источнику, галлюцинированному LLM, например программный пакет, веб-сайт или адрес электронной почты. Такие галлюцинации могут принимать форму имён пакетов, команд, URL-адресов, названий компаний или адресов электронной почты, которые направляют жертву к подконтрольной злоумышленнику сущности. Когда жертва взаимодействует с этой сущностью, атака может продолжиться.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник мог загрузить вредоносный пакет под галлюцинированным именем в PyPI или другие реестры пакетов. На практике исследователи загрузили в PyPI пустой пакет, чтобы отслеживать скачивания.</p></a>
</div>
