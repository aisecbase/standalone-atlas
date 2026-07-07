---
atlas_id: AML.T0085
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут использовать свой доступ к ИИ-сервисам организации-жертвы для сбора проприетарной или иной чувствительной информации. По мере того как организации внедряют генеративный ИИ в централизованные...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Data from AI Services
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Данные из ИИ-сервисов
url: /techniques/AML.T0085/
---

Злоумышленники могут использовать свой доступ к ИИ-сервисам организации-жертвы для сбора проприетарной или иной чувствительной информации. По мере того как организации внедряют генеративный ИИ в централизованные сервисы для доступа к данным организации, например чат-агентов, которые могут обращаться к базам данных для RAG (генерации, дополненной извлечением) и другим источникам данных через инструменты, такие сервисы становятся все более ценными целями для злоумышленников.

ИИ-агенты могут быть настроены с доступом к инструментам и источникам данных, которые недоступны пользователям напрямую. Злоумышленники могут злоупотреблять этим для сбора данных, к которым обычный пользователь не смог бы получить прямой доступ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа для привилегированных ИИ-агентов может ограничить способность злоумышленника собирать данные из ИИ-сервисов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями, унаследованными от пользователя, может ограничить способность злоумышленника собирать данные из ИИ-сервисов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0028/"><span class="relation-id">AML.M0028</span><strong>Настройка разрешений инструментов ИИ-агента</strong><p>Настройка инструментов ИИ-агента с контролем доступа, унаследованным от пользователя или вызывающего их ИИ-агента, может ограничить доступ злоумышленника к чувствительным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных из ИИ-сервисов.</p></a>
</div>
