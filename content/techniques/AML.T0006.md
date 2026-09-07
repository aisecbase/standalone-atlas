---
atlas_id: AML.T0006
atlas_type: technique
attack_ref_id: T1595
attack_ref_url: https://attack.mitre.org/techniques/T1595/
created_date: "2021-05-13"
description: Злоумышленник может зондировать или сканировать систему жертвы, чтобы собрать информацию для выбора целей. Это отличается от других техник разведки, которые не предполагают прямого взаимодействия с системой жертвы....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 7
source_name: Active Scanning
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Активное сканирование
url: /techniques/AML.T0006/
---

Злоумышленник может зондировать или сканировать систему жертвы, чтобы собрать информацию для выбора целей. Это отличается от других техник разведки, которые не предполагают прямого взаимодействия с системой жертвы.

Злоумышленники могут сканировать сеть потенциальной жертвы на наличие открытых портов, что может указывать на конкретные сервисы или инструменты, используемые жертвой. Это может включать сканирование инструментов, связанных с AI DevOps, или самих ИИ-сервисов, например публичных ИИ-чат-агентов (например, [Copilot Studio Hunter](https://github.com/mbrg/power-pwn/wiki/Modules:-Copilot-Studio-Hunter-%E2%80%90-Enum)). Они также могут отправлять письма на сервисные адреса организации и анализировать ответы на признаки того, что почтовым ящиком управляет ИИ-агент.

Информация, полученная с помощью активного сканирования, может выявить цели, которые создают возможности для других форм разведки, таких как [поиск в открытых технических базах данных](/techniques/AML.T0000), [поиск открытых материалов по анализу уязвимостей ИИ](/techniques/AML.T0001) или [сбор целей, индексируемых RAG](/techniques/AML.T0064).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Требуйте аутентификации для доступа к ИИ-эндпоинтам в продакшене и отслеживайте запросы, чтобы ограничить зондирование доступных извне ИИ-сервисов без аутентификации.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментируйте компоненты ИИ-агента так, чтобы доступный извне сервис не раскрывал сведения о других внутренних компонентах и не открывал к ним сетевой доступ.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0002 Разведка</span><p>Злоумышленники могут сканировать публичные IP-адреса, чтобы найти системы, на которых потенциально доступны панели управления Ray. По умолчанию панели Ray работают на всех сетевых интерфейсах, поэтому без дополнительных защитных механизмов они могут оказаться доступными из интернета.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0002 Разведка</span><p>Исследователи ищут на сайте целевой организации адреса электронной почты службы поддержки, которые могут обслуживаться ИИ-агентом. Затем они проверяют систему: отправляют письма и ищут в автоматических ответах признаки работы ИИ-агента.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи напрямую тестировали интерфейсы Gemini, чтобы понять, как система выбирает и задействует агентов.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0002 Разведка</span><p>Джейлбрейкнутый агент Claude от GTG-1002 сканировал диапазоны IP-адресов, связанные с целевой организацией, и её инфраструктуру на наличие уязвимостей. По результатам сканирования он выявил доступные из интернета сервисы и эндпоинты, обнаружил потенциальные уязвимости и выбрал для дальнейшего исследования уязвимость SSRF в неназванном приложении, доступном из интернета. Из опубликованных материалов нельзя установить, была ли эта уязвимость известна ранее и был ли ей присвоен идентификатор CVE.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek ran the public Langflow scanner and identified a target running Langflow 1.3.4.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek sampled approximately 100 Chinese addresses, probed roughly 40 unique systems, identified three running affected versions, inspected form endpoints, and launched parallel scanning against more than 50 remaining targets.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0002 Разведка</span><p>The framework probed primary government applications and APIs for exposed interfaces, authentication behavior, misconfigurations, and vulnerabilities. This scanning identified multiple potential paths into the targeted systems.</p></a>
</div>
