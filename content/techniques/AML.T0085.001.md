---
atlas_id: AML.T0085.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут через промпт заставить ИИ-сервис вызвать различные инструменты, к которым имеет доступ ИИ-агент. Инструменты могут извлекать данные из разных API или сервисов организации.
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 6
source_name: AI Agent Tools
subtechnique_count: 0
subtechnique_of: AML.T0085
tactics:
    - AML.TA0009
title: Инструменты ИИ-агента
url: /techniques/AML.T0085.001/
---

Злоумышленники могут через промпт заставить ИИ-сервис вызвать различные инструменты, к которым имеет доступ ИИ-агент. Инструменты могут извлекать данные из разных API или сервисов организации.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0085/"><span class="relation-id">AML.T0085</span><strong>Данные из ИИ-сервисов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа для привилегированных ИИ-агентов может ограничить способность злоумышленника собирать данные через вызов инструментов агента при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями, унаследованными от пользователя, может ограничить способность злоумышленника собирать данные через вызов инструментов агента при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0028/"><span class="relation-id">AML.M0028</span><strong>Настройка разрешений инструментов ИИ-агента</strong><p>Настройка инструментов ИИ-агента с контролем доступа, унаследованным от пользователя или вызывающего их ИИ-агента, может ограничить доступ злоумышленника к чувствительным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0009 Сбор материалов</span><p>Промпт просит агента получить все записи Salesforce с помощью инструмента `get-records`. Агент извлекает все записи из CRM организации-жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0009 Сбор материалов</span><p>Расширение Workspace находило документ и помещало его содержимое в контекст чата.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносный промпт предписывал собрать все сведения из других задач. Это вызывало инструмент Atlassian MCP, который мог обращаться к Jira-тикетам и собирать их.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносный промпт заставил Gemini с помощью подключённых инструментов ИИ-агента извлечь заголовки событий Calendar или темы писем Gmail и поместить полученные данные в URL, подконтрольные злоумышленнику.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0009 Сбор материалов</span><p>ChatGPT использовал инструменты подключённых сервисов для извлечения информации, доступной агенту жертвы. В ходе демонстрации были собраны содержимое почтового ящика и контактные данные из электронной почты.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0009 Сбор материалов</span><p>ChatGPT выполнил поиск в почтовом ящике жертвы и собрал адреса электронной почты, принадлежавшие потенциальным дополнительным целям.</p></a>
</div>
