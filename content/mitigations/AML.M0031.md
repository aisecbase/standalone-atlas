---
atlas_id: AML.M0031
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2025-10-29"
description: Усиление защиты памяти включает разработку границ доверия и безопасных процессов для того, как ИИ-агент хранит память и контекст и получает к ним доступ. Это может быть реализовано с помощью комбинации стратегий,...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - AI Model Engineering
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Memory Hardening
technique_count: 2
title: Усиление защиты памяти
url: /mitigations/AML.M0031/
---

Усиление защиты памяти включает разработку границ доверия и безопасных процессов для того, как ИИ-агент хранит память и контекст и получает к ним доступ. Это может быть реализовано с помощью комбинации стратегий, включая ограничение способности агента сохранять записи памяти за счет обязательной внешней аутентификации и проверки при обновлении памяти, выполнение семантических проверок целостности извлеченных записей памяти до выполнения агентом действий, а также внедрение механизмов контроля для мониторинга памяти и процедур устранения последствий отравления памяти.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080/"><span class="relation-id">AML.T0080</span><strong>Отравление контекста ИИ-агента</strong><p>Memory hardening reduces persistent context poisoning by controlling what an agent may save as memory, preventing saved data from becoming higher-authority instructions, and enabling poisoned records to be identified, quarantined, and rolled back.</p></a>
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><p>Memory hardening reduces persistent context poisoning by controlling what an agent may save as memory, preventing saved data from becoming higher-authority instructions, and enabling poisoned records to be identified, quarantined, and rolled back.</p></a>
</div>
