---
atlas_id: AML.TA0012
atlas_type: tactic
attack_ref_id: TA0004
attack_ref_url: https://attack.mitre.org/tactics/TA0004/
created_date: "2023-10-25"
description: Злоумышленник пытается получить более высокий уровень прав. Повышение привилегий включает техники, которые злоумышленники используют для получения более высокого уровня прав в системе или сети. Злоумышленники часто...
generated: true
generated_by: atlasgen
modified_date: "2023-10-25"
procedure_count: 14
source_name: Privilege Escalation
technique_count: 4
title: Повышение привилегий
url: /tactics/AML.TA0012/
---

Злоумышленник пытается получить более высокий уровень прав.

Повышение привилегий включает техники, которые злоумышленники используют для получения более высокого уровня прав в системе или сети.
Злоумышленники часто могут попасть в сеть и исследовать ее с непривилегированным доступом, но для достижения своих целей им требуются повышенные права.
Распространенные подходы включают эксплуатацию слабых мест системы, ошибок конфигурации и уязвимостей.
Примеры повышенного доступа включают:

- уровень SYSTEM/root
- локальный администратор
- учетная запись пользователя с правами, близкими к административным
- учетные записи пользователей с доступом к конкретной системе или возможностью выполнять конкретную функцию

Эти техники часто пересекаются с техниками закрепления, поскольку функции ОС, которые позволяют злоумышленнику закрепиться, могут выполняться в повышенном контексте.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0012/"><span class="relation-id">AML.T0012</span><strong>Действующие учетные записи</strong></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0054/"><span class="relation-id">AML.T0054</span><strong>Джейлбрейк LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0105/"><span class="relation-id">AML.T0105</span><strong>Выход на хост</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0012 Повышение привилегий</span><p>Кроме того, промпт может заставить LLM вызывать плагины, которые пользователь не запрашивал. В этом примере исследователь показал, как плагин `WebPilot` обращается к плагину `Expedia`.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователи скомпрометировали плагин `search_enterprise`, заставив LLM переопределить часть поведения и использовать в ответе только извлеченный объект `EmailMessage`.</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0012 Повышение привилегий</span><p>Скомпрометированные учетные данные дали злоумышленникам доступ к облачным средам, где были размещены сервисы больших языковых моделей (LLM).</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0012 Повышение привилегий</span><p>При следующем взаимодействии жертвы с Gemini вызывалось расширение Workspace.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт запрашивал информацию, доступную ИИ-агенту через инструменты Atlassian MCP. Это приводило к вызову этих инструментов через MCP и повышало привилегии исследователей в JSM-экземпляре жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Промпт-инъекция задействовала возможность Cursor вызывать команды через инструмент `run_terminal_cmd`. Перед выполнением команды оболочки Cursor запросил подтверждение пользователя, что потенциально снижало риск атаки.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учетные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователь отправил ClawdBot промпт `root`; в ответ ClawdBot вызвал навык `bash`, запущенный от имени пользователя root.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0012 Повышение привилегий</span><p>Claude Code выполнил команду оболочки через свой инструмент `bash`.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный скрипт использовал похищенный Gateway-токен для аутентификации, что позволяло затем выполнять вызовы к OpenClaw Gateway API в системе жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный скрипт отключал песочницу OpenClaw, заставляя агента выполнять команды напрямую на хост-машине, а не внутри Docker-контейнера. Для этого в OpenClaw Gateway API отправлялся запрос `config.patch`, устанавливающий `tools.exec.host` в значение `&#34;gateway&#34;`.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0012 Повышение привилегий</span><p>Промпты исследователей вызывали инструменты ИИ-агента и были нацелены на цепочки вызовов, которые могут привести к выполнению кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователи включили техники выхода из изоляции, предназначенные для обхода ограничений, которые песочница может накладывать на выполнение кода.</p></a>
</div>


Показано 12 из 14 примеров.
