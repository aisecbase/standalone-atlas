---
atlas_id: AML.T0086
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Инструменты ИИ-агента, способные выполнять операции записи, могут вызываться для эксфильтрации данных злоумышленнику. Чувствительная информация может быть закодирована во входных параметрах инструмента и передана в...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 8
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 10
source_name: Exfiltration via AI Agent Tool Invocation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Эксфильтрация через вызов инструмента ИИ-агента
url: /techniques/AML.T0086/
---

Инструменты ИИ-агента, способные выполнять операции записи, могут вызываться для эксфильтрации данных злоумышленнику. Чувствительная информация может быть закодирована во входных параметрах инструмента и передана в контролируемое злоумышленником место назначения, например в почтовый ящик, документ или на сервер, как часть действия, выглядящего легитимным. Варианты включают отправку электронных писем, создание или изменение документов, обновление записей CRM или даже генерацию медиа, таких как изображения или видео.

Сам вызываемый инструмент может быть легитимным, но при этом вызван злоумышленником через [промпт-инъекцию LLM](/techniques/AML.T0051), либо инструмент может быть вредоносным (см. [Отравление инструмента ИИ-агента](/techniques/AML.T0110)).

[Отравление инструмента ИИ-агента](/techniques/AML.T0110) также может использоваться для манипуляции входными данными и местом назначения отдельного легитимного инструмента, вызываемого в ходе обычного использования жертвой.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа к использованию инструментов привилегированными ИИ-агентами может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями на использование инструментов, унаследованными от пользователя, может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0028/"><span class="relation-id">AML.M0028</span><strong>Настройка разрешений инструментов ИИ-агента</strong><p>Настройка инструментов ИИ-агента с контролем доступа, унаследованным от пользователя или вызывающего их ИИ-агента, может ограничить возможности злоумышленника в системе, включая злоупотребление вызовами инструментов и эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0029/"><span class="relation-id">AML.M0029</span><strong>Участие человека в действиях ИИ-агента</strong><p>Требование подтверждения пользователем вызовов инструментов ИИ-агента может предотвратить автоматическое выполнение инструментов злоумышленником.</p></a>
<a class="relation-item" href="/mitigations/AML.M0030/"><span class="relation-id">AML.M0030</span><strong>Ограничение вызова инструментов ИИ-агента при работе с недоверенными данными</strong><p>Ограничение автоматического использования инструментов при наличии недоверенных данных может помешать злоумышленникам вызывать инструменты через промпт-инъекции.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для компрометации источников чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для компрометации источников чувствительных данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0010 Эксфильтрация</span><p>Промпт просит агента отправить результаты письмом на адрес, выбранный исследователями, с помощью почтового инструмента агента. Исследователи успешно эксфильтруют целевые данные через вызов инструмента.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносный промпт предписывал опубликовать собранные сведения о тикетах в ответе к обращению. Это вызывало инструмент Atlassian MCP, который выполнял запрошенное действие и эксфильтровал данные в место, доступное исследователям на портале JSM.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Файлы с учетными данными были эксфильтрированы на сервер исследователя командой `curl`, вызванной через инструмент Cursor `run_terminal_cmd`.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0010 Эксфильтрация</span><p>Когда организации отправляли письма через инструмент `postmark-mcp`, все содержимое этих писем эксфильтровывалось злоумышленнику через адрес, добавленный в поле скрытой копии (BCC).</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0010 Эксфильтрация</span><p>Промпт инструктировал ИИ-агента сохранить файлы с учетными данными в лишнем параметре MCP-инструмента, чтобы эксфильтровать их через MCP-соединение.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0010 Эксфильтрация</span><p>Собранная информация о хосте жертвы эксфильтрировалась, когда ИИ-сервис следовал инструкциям получить URL на подконтрольном злоумышленнику домене с данными, встроенными в параметр запроса.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Gemini открыл сформированный URL с помощью Android Utilities, тем самым передав на подконтрольный злоумышленнику сервер данные жертвы из Calendar или электронной почты.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0010 Эксфильтрация</span><p>Скомпрометированный ИИ-агент вызывает инструмент с сетевым доступом или возможностью записи, чтобы передать чувствительную информацию в место назначения, подконтрольное злоумышленнику. Затем агент может продолжить и завершить легитимную задачу пользователя, тем самым скрывая неавторизованную передачу.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: Data Exfiltration Attack on ChatGPT</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0010 Эксфильтрация</span><p>ChatGPT invoked its URL-opening capability to request the selected adversary-controlled URLs. The researchers reconstructed the sensitive data from the character and position encoded in the resulting server requests. The requests originated from OpenAI&#39;s infrastructure rather than the victim&#39;s endpoint or corporate network.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Claude Code GitHub Action Secret Exposure</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>The researchers could use WebFetch, Bash, GitHub MCP, and Action logs as potential exfiltration channels depending on the tools available in the workflow configuration.</p></a>
</div>
