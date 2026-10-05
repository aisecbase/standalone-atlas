---
atlas_id: AML.T0006.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Злоумышленники могут обращаться к документированным или недокументированным API SaaS-платформ для размещения ИИ-агентов, чтобы выявлять цели для атак на ИИ-агентов. SaaS-платформы и платформы ИИ-агентов могут...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 0
source_name: Query Platform Metadata APIs
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Запросы к API метаданных платформ
url: /techniques/AML.T0006.001/
---

Злоумышленники могут обращаться к документированным или недокументированным API SaaS-платформ для размещения ИИ-агентов, чтобы выявлять цели для атак на ИИ-агентов. SaaS-платформы и платформы ИИ-агентов могут открывать доступ к API поставщика для управления платформой или работы со службами идентификации. Эти API возвращают идентификаторы тенантов, сред или развёртываний, в том числе для ресурсов с ошибками конфигурации или ресурсов, непреднамеренно доступных пользователям без аутентификации.

Зафиксированы случаи, когда злоумышленники использовали такую функциональность для сбора информации о SaaS-платформах, например через OSINT-страницу AADInternals,[[aadinternals]] — онлайн-инструмент OSINT, демонстрирующий метод разведки Entra ID с помощью недокументированного API. Этот недокументированный API Power Platform мог использоваться для выявления идентификаторов сред, которые затем могут применяться при сканировании для поиска общедоступных агентов. После выявления такого злоупотребления в инструменте ввели обязательную аутентификацию.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006.000/"><span class="relation-id">AML.T0006.000</span><strong>Получение списка размещённых ресурсов ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Сканирование для поиска доступной из интернета инфраструктуры ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Зондирование каналов запуска ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>ИИ-ханипоты</strong><p>Фиксируя действия злоумышленников по распознаванию систем и сервисов, ханипоты могут помогать на раннем этапе сигнализировать о массовом поиске доступных извне ИИ-целей и выявлять поверхности атак, которые исследуют злоумышленники.</p></a>
</div>
