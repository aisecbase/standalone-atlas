---
atlas_id: AML.T0006.002
atlas_type: technique
attack_ref_id: T1595
attack_ref_url: https://attack.mitre.org/techniques/T1595/
created_date: "2026-09-15"
description: Злоумышленники могут сканировать сетевые порты и сервисы, чтобы выявлять развёрнутые серверные компоненты ИИ, конечные точки сервисов инференса моделей и инфраструктуру ИИ-агентов, доступные из интернета....
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
source_name: Scan for Exposed AI Infrastructure
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Сканирование для поиска доступной из интернета инфраструктуры ИИ
url: /techniques/AML.T0006.002/
---

Злоумышленники могут сканировать сетевые порты и сервисы, чтобы выявлять развёрнутые серверные компоненты ИИ, конечные точки сервисов инференса моделей и инфраструктуру ИИ-агентов, доступные из интернета. Самостоятельно развёрнутые среды выполнения ИИ обычно принимают соединения на предсказуемых, широко известных портах и предоставляют стандартные пути API. Это позволяет злоумышленникам эффективно находить хосты для дальнейшей проверки и предполагать, какая платформа на них работает.

Найдя хосты для дальнейшей проверки, злоумышленники могут напрямую взаимодействовать с ними, чтобы подтвердить наличие работающих ИИ-сервисов, распознать используемый программный стек и получить сведения о версиях и конфигурации. Эта информация может использоваться для выбора целей, уязвимости которых можно эксплуатировать, и адаптации последующих попыток доступа.


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
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Запросы к API метаданных платформ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Зондирование каналов запуска ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>ИИ-ханипоты</strong><p>Фиксируя действия злоумышленников по распознаванию систем и сервисов, ханипоты могут помогать на раннем этапе сигнализировать о массовом поиске доступных извне ИИ-целей и выявлять поверхности атак, которые исследуют злоумышленники.</p></a>
</div>
