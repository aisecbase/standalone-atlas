---
atlas_id: AML.T0069
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленник пытается выявить системную информацию большой языковой модели (LLM). Эта информация может находиться в конфигурационном файле, содержащем системные инструкции, или извлекаться через взаимодействие с LLM....
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Discover LLM System Information
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление системной информации LLM
url: /techniques/AML.T0069/
---

Злоумышленник пытается выявить системную информацию большой языковой модели (LLM). Эта информация может находиться в конфигурационном файле, содержащем системные инструкции, или извлекаться через взаимодействие с LLM. Интересующая информация может включать полный системный промпт, специальные символы, значимые для LLM, или ключевые слова, указывающие на функциональность, доступную LLM. Сведения о том, как LLM проинструктирована, могут использоваться злоумышленником, чтобы понять возможности системы и упростить создание вредоносных промптов.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><span class="relation-meta">Подтехника</span></a>
</div>
