---
atlas_id: AML.T0067
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут использовать промпты к большой языковой модели (LLM), которые манипулируют различными компонентами ее ответа, чтобы он выглядел заслуживающим доверия для пользователя. Это помогает злоумышленнику...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: LLM Trusted Output Components Manipulation
subtechnique_count: 1
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Манипуляция доверенными компонентами ответа LLM
url: /techniques/AML.T0067/
---

Злоумышленники могут использовать промпты к большой языковой модели (LLM), которые манипулируют различными компонентами ее ответа, чтобы он выглядел заслуживающим доверия для пользователя. Это помогает злоумышленнику продолжать действовать в среде жертвы и избегать обнаружения пользователями, с которыми он взаимодействует.

LLM может быть проинструктирована адаптировать формулировки ответа так, чтобы он выглядел для пользователя более заслуживающим доверия, или попытаться склонить пользователя к определенным действиям. К другим компонентам ответа, которыми можно манипулировать, относятся ссылки, рекомендуемые последующие действия, метаданные извлеченных документов и [ссылки на источники](/techniques/AML.T0067.000).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0067.000/"><span class="relation-id">AML.T0067.000</span><strong>Ссылки на источники</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт предписывал ИИ-ассистенту для программирования не упоминать изменения кода в ответах, чтобы не вызывать подозрений у жертвы и не оставлять следов в журналах ассистента. Это позволяет вредоносному файлу правил скрытно распространяться по кодовой базе без следов в истории или журналах, которые могли бы помочь командам безопасности обнаружить проблему.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вывод был изменен так, чтобы избежать очевидной атрибуции и использовать ссылки или изображения Markdown в справочном стиле, обходившие сокрытие ссылок.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Poisoned GGUF Templates: Inference-Time Supply Chain Attack</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0007 Уклонение от защиты</span><p>The injected instruction causes the model to include attacker-selected links, references, or other response components in a form that appears relevant or trustworthy to the user.</p></a>
</div>
