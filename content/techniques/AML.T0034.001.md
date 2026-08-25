---
atlas_id: AML.T0034.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут создавать входные данные, специально рассчитанные на увеличение объёма вычислительных ресурсов, необходимых для обработки. Для моделей генеративного ИИ злоумышленники могут использовать длинные...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Resource-Intensive Queries
subtechnique_count: 0
subtechnique_of: AML.T0034
tactics:
    - AML.TA0011
title: Ресурсоёмкие запросы
url: /techniques/AML.T0034.001/
---

Злоумышленники могут создавать входные данные, специально рассчитанные на увеличение объёма вычислительных ресурсов, необходимых для обработки.

Для моделей генеративного ИИ злоумышленники могут использовать длинные входные последовательности, запросы на крайне длинные ответы или промпты, требующие сложных рассуждений, чтобы увеличить вычислительные затраты [[genai]]. Для vision- и language-моделей могут применяться «sponge examples» [[arxiv]], которые максимизируют энергопотребление и задержку принятия решения. Использование небольшого числа вычислительно затратных запросов вместо простого потока чрезмерных запросов может быть сложнее обнаружить, заблокировать или ограничить.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0034.000/"><span class="relation-id">AML.T0034.000</span><strong>Чрезмерные запросы</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0034.002/"><span class="relation-id">AML.T0034.002</span><strong>Потребление ресурсов агентом</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Submit controlled requests designed to consume disproportionate inference resources. Verify input constraints, timeouts, workload limits, resource isolation, and cost monitoring.</p></a>
<a class="relation-item" href="/mitigations/AML.M0036/"><span class="relation-id">AML.M0036</span><strong>Limit AI Workload Resource Consumption</strong><p>Bound input size, output size, execution time, memory, and compute consumed by resource-intensive queries.</p></a>
</div>


## Источники

- [\[2006.03463\] Sponge Examples: Energy-Latency Attacks on Neural Networks](https://arxiv.org/abs/2006.03463)
- [OWASP Top 10 for LLM Applications 2025](https://genai.owasp.org/resource/owasp-top-10-for-llm-applications-2025/)
