---
atlas_id: AML.T0018.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may modify templates, role delimiters, embedded system instructions, tokenizer settings, tool-call formatting, or other artifact-bundled logic that constructs the context sent to an AI model. Model file...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Modify Prompt Construction Logic
subtechnique_count: 0
subtechnique_of: AML.T0018
tactics:
    - AML.TA0001
    - AML.TA0006
title: Modify Prompt Construction Logic
url: /techniques/AML.T0018.003/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may modify templates, role delimiters, embedded system instructions, tokenizer settings, tool-call formatting, or other artifact-bundled logic that constructs the context sent to an AI model. Model file formats such as GGUF can package this logic alongside model weights in a single distributable artifact. A compatible inference runtime may interpret the modified logic during future inference requests, enabling persistent covert instruction injection, altered instruction precedence, redirected tool use, or manipulated model output without changing model weights.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Poisoned GGUF Templates: Inference-Time Supply Chain Attack</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The adversary modifies the chat template bundled with the model artifact. The modified template injects attacker-controlled instructions into the model context when its trigger is present, while leaving the model weights unchanged.</p></a>
</div>
