---
atlas_id: AML.M0036
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-07-31"
description: Limit the resources that an AI request, inference job, or agent workflow can consume. Set bounds on input size, batch size, execution time, memory, compute, and output size. Generative AI services should also limit...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Limit AI Workload Resource Consumption
technique_count: 4
title: Limit AI Workload Resource Consumption
url: /mitigations/AML.M0036/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Limit the resources that an AI request, inference job, or agent workflow can consume. Set bounds on input size, batch size, execution time, memory, compute, and output size. Generative AI services should also limit context and output tokens. Agentic AI systems should limit iterations, retries, tool calls, parallel tasks, delegation depth, and downstream spending.

Apply resource limits across the complete workflow because one request may initiate multiple model calls or external actions. Use timeouts, cost ceilings, circuit breakers, and safe termination conditions to prevent individual workloads from exhausting shared resources. These controls complement query-rate limits, which address large numbers of otherwise inexpensive requests.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Limit the resources consumed by individual requests to reduce denial of service from computationally expensive inputs.</p></a>
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong><p>Apply resource budgets to AI requests and workflows to limit costs from resource-intensive queries, excessive output, and uncontrolled agentic activity.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.001/"><span class="relation-id">AML.T0034.001</span><strong>Ресурсоёмкие запросы</strong><p>Bound input size, output size, execution time, memory, and compute consumed by resource-intensive queries.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.002/"><span class="relation-id">AML.T0034.002</span><strong>Потребление ресурсов агентом</strong><p>Limit agent iterations, tool calls, fan-out, runtime, and downstream spending to constrain agentic resource consumption.</p></a>
</div>
