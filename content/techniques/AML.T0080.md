---
atlas_id: AML.T0080
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут пытаться манипулировать контекстом, который использует большая языковая модель (LLM) ИИ-агента, чтобы повлиять на генерируемые ею ответы или выполняемые действия. Это позволяет злоумышленнику...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: AI Agent Context Poisoning
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Отравление контекста ИИ-агента
url: /techniques/AML.T0080/
---

Злоумышленники могут пытаться манипулировать контекстом, который использует большая языковая модель (LLM) ИИ-агента, чтобы повлиять на генерируемые ею ответы или выполняемые действия. Это позволяет злоумышленнику устойчиво изменить поведение целевого агента и продвинуться к своим целям.

Отравление контекста может выполняться через промпт, который заставляет LLM добавить инструкции или предпочтения в память (см. [Память](/techniques/AML.T0080.000)), либо через обычный промпт к LLM, использующей предыдущие сообщения в цепочке как часть своего контекста (см. [Цепочка сообщений](/techniques/AML.T0080.001)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0080.001/"><span class="relation-id">AML.T0080.001</span><strong>Цепочка сообщений</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0031/"><span class="relation-id">AML.M0031</span><strong>Усиление защиты памяти</strong><p>Memory hardening reduces persistent context poisoning by controlling what an agent may save as memory, preventing saved data from becoming higher-authority instructions, and enabling poisoned records to be identified, quarantined, and rolled back.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>AI Red Team</strong><p>Attempt to persist malicious instructions in agent memory and long-lived threads. Verify authorization for context changes, integrity checks, trust labeling, expiration, user visibility, and remediation.</p></a>
</div>
