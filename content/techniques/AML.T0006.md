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
procedure_count: 3
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
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Require authenticated access to production AI endpoints and monitor queries to limit unauthenticated probing of exposed AI services.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Segment AI agent components so an exposed service does not reveal or provide reachability to additional internal components.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0002 Разведка</span><p>Злоумышленники могут сканировать публичные IP-адреса, чтобы найти системы, на которых потенциально доступны панели управления Ray. По умолчанию панели Ray работают на всех сетевых интерфейсах, поэтому без дополнительных защитных механизмов они могут оказаться доступными из интернета.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0002 Разведка</span><p>Исследователи ищут на сайте целевой организации адреса электронной почты службы поддержки, которые могут обслуживаться ИИ-агентом. Затем они проверяют систему: отправляют письма и ищут в автоматических ответах признаки работы ИИ-агента.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0002 Разведка</span><p>The researchers directly probed Gemini interfaces to understand its agent selection and execution behavior.</p></a>
</div>
