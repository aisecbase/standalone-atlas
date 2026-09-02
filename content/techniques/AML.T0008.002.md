---
atlas_id: AML.T0008.002
atlas_type: technique
attack_ref_id: T1583.001
attack_ref_url: https://attack.mitre.org/techniques/T1583/001/
created_date: "2025-03-12"
description: Злоумышленники могут приобретать домены, которые можно использовать при выборе целей. Доменные имена — это понятные человеку имена, представляющие один или несколько IP-адресов. Их можно купить или в некоторых случаях...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Domains
subtechnique_count: 0
subtechnique_of: AML.T0008
tactics:
    - AML.TA0003
title: Домены
url: /techniques/AML.T0008.002/
---

Злоумышленники могут приобретать домены, которые можно использовать при выборе целей. Доменные имена — это понятные человеку имена, представляющие один или несколько IP-адресов. Их можно купить или в некоторых случаях получить бесплатно.

Злоумышленники могут использовать приобретённые домены для различных целей (см. [ATT&CK](https://attack.mitre.org/techniques/T1583/001/)). Крупные наборы данных ИИ часто распространяются в виде списков URL, указывающих на отдельные элементы данных. Злоумышленники могут приобретать домены с истёкшим сроком регистрации, URL которых входят в такие наборы данных, и заменять отдельные элементы данных отравленными примерами ([Публикация отравленных ИИ-артефактов: наборы данных](/techniques/AML.T0115.000)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0008/"><span class="relation-id">AML.T0008</span><strong>Получение инфраструктуры</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0008.000/"><span class="relation-id">AML.T0008.000</span><strong>Рабочие пространства для разработки ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи находят в датасете домены с истекшей регистрацией и выкупают их.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь зарегистрировал домен `clawdhub-skill.com` для своего веб-сервера.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи зарегистрировали домен и развернули HTTPS-сайт, который выполнял роль конечной точки ретрансляции.</p></a>
</div>
