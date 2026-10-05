---
atlas_id: AML.T0006
atlas_type: technique
attack_ref_id: T1595
attack_ref_url: https://attack.mitre.org/techniques/T1595/
created_date: "2021-05-13"
description: Злоумышленники могут активно сканировать доступные из интернета ИИ-системы и ресурсы для выбора целей. Они могут выявлять ресурсы или системы с ошибками конфигурации либо использующие версии с известными уязвимостями....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-09-15"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: Active Scanning
subtechnique_count: 4
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Активное сканирование
url: /techniques/AML.T0006/
---

Злоумышленники могут активно сканировать доступные из интернета ИИ-системы и ресурсы для выбора целей. Они могут выявлять ресурсы или системы с ошибками конфигурации либо использующие версии с известными уязвимостями. Это отличается от других техник разведки, которые не предполагают прямого взаимодействия с системой жертвы.

Поскольку ИИ-системы часто развёртываются в облачных средах, злоумышленники могут использовать различные методы получения списков ресурсов, чтобы эффективно выявлять ИИ-системы и получать сведения о платформах, обеспечивающих их работу. Для систем ИИ-агентов и систем, размещённых на SaaS-платформах, понятие системы жертвы может охватывать общую платформу, на которой развёрнуты агенты жертвы, и средства управления этой платформой на стороне поставщика.

Злоумышленники могут зондировать систему жертвы, чтобы собирать информацию и метаданные для выбора целей, либо активно сканировать сеть потенциальной жертвы или доступные развёрнутые серверы на наличие определённых открытых портов, чтобы находить различные типы развёрнутых ИИ-сервисов. Эти методы также могут включать непосредственное сканирование ресурсов жертвы для поиска общедоступных ИИ-агентов или доступных чат-интерфейсов ИИ-агентов на SaaS-платформах (например, [Copilot Studio Hunter](https://github.com/mbrg/power-pwn/wiki/Modules:-Copilot-Studio-Hunter-%E2%80%90-Enum)).

Сведения, полученные с помощью активного сканирования, могут помочь выявить цели, открывающие возможности для других форм разведки, таких как [поиск в открытых технических базах данных](/techniques/AML.T0000), [поиск на открытых сайтах и доменах](/techniques/AML.T0095), [поиск открытых материалов по анализу уязвимостей ИИ](/techniques/AML.T0001) или [сбор целей, индексируемых RAG](/techniques/AML.T0064).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006.000/"><span class="relation-id">AML.T0006.000</span><strong>Получение списка размещённых ресурсов ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Запросы к API метаданных платформ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Сканирование для поиска доступной из интернета инфраструктуры ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Зондирование каналов запуска ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Требуйте аутентификации для доступа к ИИ-эндпоинтам в продакшене и отслеживайте запросы, чтобы ограничить зондирование доступных извне ИИ-сервисов без аутентификации.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментируйте компоненты ИИ-агента так, чтобы доступный извне сервис не раскрывал сведения о других внутренних компонентах и не открывал к ним сетевой доступ.</p></a>
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>ИИ-ханипоты</strong><p>Фиксируя действия злоумышленников по распознаванию систем и сервисов, ханипоты могут помогать на раннем этапе сигнализировать о массовом поиске доступных извне ИИ-целей и выявлять поверхности атак, которые исследуют злоумышленники.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0002 Разведка</span><p>Злоумышленники могут сканировать публичные IP-адреса, чтобы найти системы, на которых потенциально доступны панели управления Ray. По умолчанию панели Ray работают на всех сетевых интерфейсах, поэтому без дополнительных защитных механизмов они могут оказаться доступными из интернета.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи напрямую тестировали интерфейсы Gemini, чтобы понять, как система выбирает и задействует агентов.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0002 Разведка</span><p>Джейлбрейкнутый агент Claude от GTG-1002 сканировал диапазоны IP-адресов, связанные с целевой организацией, и её инфраструктуру на наличие уязвимостей. По результатам сканирования он выявил доступные из интернета сервисы и эндпоинты, обнаружил потенциальные уязвимости и выбрал для дальнейшего исследования уязвимость SSRF в неназванном приложении, доступном из интернета. Из опубликованных материалов нельзя установить, была ли эта уязвимость известна ранее и был ли ей присвоен идентификатор CVE.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek запустил общедоступный сканер Langflow и выявил целевую систему, на которой работал Langflow версии 1.3.4.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek отобрал около 100 адресов в Китае, прозондировал примерно 40 различных систем, обнаружил три системы с затронутыми версиями, проверил эндпоинты форм и запустил параллельное сканирование более чем 50 оставшихся целей.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Мультиагентный фреймворк скомпрометировал государственные системы Тайваня</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0002 Разведка</span><p>Фреймворк зондировал основные государственные приложения и API, чтобы выявить доступные извне интерфейсы, особенности аутентификации, ошибки конфигурации и уязвимости. Это сканирование выявило несколько возможных путей проникновения в целевые системы.</p></a>
</div>
