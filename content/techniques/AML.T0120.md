---
atlas_id: AML.T0120
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут использовать репозитории ИИ-артефактов не по назначению — в качестве асинхронных каналов командования и управления. Команды, полезные нагрузки или задания могут размещаться в объектах репозитория,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: AI Artifact Repository
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0014
title: Репозиторий ИИ-артефактов
url: /techniques/AML.T0120/
---

Злоумышленники могут использовать репозитории ИИ-артефактов не по назначению — в качестве асинхронных каналов командования и управления. Команды, полезные нагрузки или задания могут размещаться в объектах репозитория, которые скомпрометированная система может извлекать или периодически опрашивать. Затем скомпрометированная система может записывать обратно в репозиторий результаты выполнения, сведения о состоянии или собранную информацию для последующего получения злоумышленником.

Для такой коммуникации могут использоваться штатные операции с артефактами и репозиториями, например чтение или обновление содержимого, метаданных либо ревизий артефакта через API или интерфейс системы контроля версий. Репозиторий выступает в роли очереди сообщений или электронного тайника (dead drop), позволяя сторонам обмениваться информацией без постоянного прямого соединения; при этом такая активность может сливаться с легитимным трафиком, связанным с артефактами.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0014/"><span class="relation-id">AML.TA0014</span><strong>Командование и управление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0014 Командование и управление</span><p>The agents used repositories they controlled as asynchronous command-and-control channels. Compromised workers retrieved staged commands or payloads and wrote results into dataset objects for retrieval through the Hugging Face API or Git.</p></a>
</div>
