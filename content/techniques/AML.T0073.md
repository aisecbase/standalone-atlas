---
atlas_id: AML.T0073
atlas_type: technique
attack_ref_id: T1656
attack_ref_url: https://attack.mitre.org/techniques/T1656/
created_date: "2025-04-14"
description: Злоумышленники могут выдавать себя за доверенное лицо или организацию, чтобы убедить цель и обманом заставить ее выполнить какое-либо действие в интересах злоумышленников. Например, злоумышленники могут...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 5
source_name: Impersonation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Имперсонация
url: /techniques/AML.T0073/
---

Злоумышленники могут выдавать себя за доверенное лицо или организацию, чтобы убедить цель и обманом заставить ее выполнить какое-либо действие в интересах злоумышленников. Например, злоумышленники могут взаимодействовать с жертвами через [фишинг](/techniques/AML.T0052) или [целевой фишинг через LLM для социальной инженерии](/techniques/AML.T0052.000), выдавая себя за известного отправителя, например руководителя, коллегу или стороннего поставщика. Сформированное доверие затем может использоваться для достижения конечных целей злоумышленника, возможно против нескольких жертв.

Злоумышленники могут нацеливаться на ресурсы, относящиеся к жизненному циклу AI DevOps, такие как репозитории моделей, реестры контейнеров и реестры программного обеспечения.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0007 Уклонение от защиты</span><p>Сотрудники целевой компании нашли поддельную организацию на Hugging Face и присоединились к ней. Поскольку имя учётной записи этой организации совпадало с названием реальной организации или выглядело почти так же, сотрудники приняли учётную запись за официальную.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Получив аутентифицированный аккаунт под личностью жертвы, исследователи успешно выдают себя за жертву и избегают обнаружения.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0007 Уклонение от защиты</span><p>Получив аутентифицированный аккаунт под личностью жертвы, злоумышленник успешно выдал себя за жертву и избежал обнаружения.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее динамические команды, сгенерированные ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0007 Уклонение от защиты</span><p>В письме злоумышленники выдавали себя за представителя государственного министерства.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленник выдал себя за Postmark, опубликовав в npm легитимную версию пакета `postmark-mcp`. Поскольку Postmark сам не зарегистрировал имя `postmark-mcp` в npm, злоумышленник смог занять это имя. Легитимных пользователей обманом заставили использовать npm-пакет, хотя им не управляли официальные разработчики `postmark-mcp`.</p></a>
</div>
