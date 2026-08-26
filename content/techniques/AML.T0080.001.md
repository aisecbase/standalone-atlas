---
atlas_id: AML.T0080.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут внедрять вредоносные инструкции в цепочку сообщений большой языковой модели (LLM), чтобы вызвать изменения поведения, сохраняющиеся до конца этой цепочки. Цепочка сообщений может продолжаться...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Thread
subtechnique_count: 0
subtechnique_of: AML.T0080
tactics:
    - AML.TA0006
title: Цепочка сообщений
url: /techniques/AML.T0080.001/
---

Злоумышленники могут внедрять вредоносные инструкции в цепочку сообщений большой языковой модели (LLM), чтобы вызвать изменения поведения, сохраняющиеся до конца этой цепочки. Цепочка сообщений может продолжаться длительное время и охватывать несколько сессий.

Вредоносные инструкции могут внедряться через прямую или косвенную промпт-инъекцию. Прямая инъекция может происходить в случаях, когда злоумышленник получил API-ключи LLM пользователя и может напрямую добавлять запросы в любую цепочку сообщений.

По мере роста лимитов токенов LLM ИИ-системы могут использовать более крупные контекстные окна, из-за чего вредоносные инструкции дольше сохраняются в цепочке сообщений.

Отравление цепочки сообщений может затрагивать нескольких пользователей, если LLM используется в сервисе с общими цепочками сообщений. Например, если агент активен в канале Slack с несколькими участниками, одно вредоносное сообщение от одного пользователя может повлиять на поведение агента в будущих взаимодействиях с другими пользователями.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080/"><span class="relation-id">AML.T0080</span><strong>Отравление контекста ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Внедряйте контролируемые вредоносные инструкции в долгоживущие цепочки сообщений или цепочки сообщений с общим доступом. Проверяйте изоляцию контекста и его обработку с учётом уровня доверия, сброс цепочек сообщений, сроки действия и мониторинг.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0006 Закрепление</span><p>Злоумышленник мог создавать вредоносные промпты, манипулирующие контекстом цепочки сообщений; этот эффект сохранялся бы до конца такой цепочки.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0006 Закрепление</span><p>Контекст всех новых диалогов был отравлен вредоносным промптом. Измененное поведение OpenClaw должно было срабатывать, когда жертва приветствовала агента.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0006 Закрепление</span><p>Внедрённые инструкции оставались в контексте диалога — в том числе инструкции из содержимого Calendar, скрытого за элементом интерфейса &#39;Show more&#39;, — чтобы влиять на последующие ходы диалога.</p></a>
</div>
