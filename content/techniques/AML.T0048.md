---
atlas_id: AML.T0048
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2022-10-27"
description: Злоумышленники могут злоупотреблять доступом к системе организации-жертвы и использовать ее ресурсы или возможности для достижения своих целей, причиняя ущерб за пределами этой системы. Такой ущерб может затрагивать...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: External Harms
subtechnique_count: 5
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Внешний ущерб
url: /techniques/AML.T0048/
---

Злоумышленники могут злоупотреблять доступом к системе организации-жертвы и использовать ее ресурсы или возможности для достижения своих целей, причиняя ущерб за пределами этой системы.

Такой ущерб может затрагивать саму организацию (например, финансовый или репутационный ущерб), ее пользователей (например, ущерб пользователям) или широкую общественность (например, общественный вред).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0048.000/"><span class="relation-id">AML.T0048.000</span><strong>Финансовый ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.001/"><span class="relation-id">AML.T0048.001</span><strong>Репутационный ущерб</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.002/"><span class="relation-id">AML.T0048.002</span><strong>Общественный вред</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.003/"><span class="relation-id">AML.T0048.003</span><strong>Ущерб пользователям</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0048.004/"><span class="relation-id">AML.T0048.004</span><strong>Кража интеллектуальной собственности ИИ</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0011 Воздействие</span><p>Эксфильтрированные данные могут включать чувствительные или частные данные, например проприетарные данные, хранящиеся в Google Drive, а также контакты и фотографии пользователя. В результате пользователю может быть причинен финансовый, репутационный и иной ущерб.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0011 Воздействие</span><p>Если модели компании будут изменены так, чтобы выдавать ложную информацию, это может привести к разным видам ущерба, включая финансовый и репутационный.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0011 Воздействие</span><p>В этом proof-of-concept исследователь лишь отправил запрос на свой сервер и предупредил пользователя об опасности использования навыков без чтения исходного кода, не причинив вреда. Однако вместо этого он мог доставить вредоносную полезную нагрузку и причинить разный ущерб: эксфильтровать кодовую базу пользователя, внедрить в нее бэкдоры, украсть учетные данные, установить вредоносное ПО или криптомайнеры либо выполнить любые другие действия, доступные Claude Code.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0011 Воздействие</span><p>Эксфильтрированные письма могли включать транзакционные письма, раскрывающие конфиденциальную информацию о клиентах организации, и маркетинговые письма, раскрывающие клиентскую базу организации.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0011 Воздействие</span><p>При эксплуатации против реального корпоративного пользователя атака могла раскрыть конфиденциальные бизнес-данные и нанести ущерб организации или затронутым пользователям.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0011 Воздействие</span><p>Атака подвергала сотрудников поддержки и клиентов рискам захвата сеанса, несанкционированного доступа к данным и потенциального выполнения вредоносного ПО, что приводило к прямому ущербу для безопасности и конфиденциальности на уровне пользователя.</p></a>
</div>
