---
atlas_id: AML.T0075
atlas_type: technique
attack_ref_id: T1526
attack_ref_url: https://attack.mitre.org/techniques/T1526/
created_date: "2025-04-14"
description: 'После получения доступа злоумышленники могут пытаться перечислить облачные сервисы, запущенные в системе. Эти методы могут различаться в зависимости от модели: platform-as-a-service (PaaS), infrastructure-as-a-service...'
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Cloud Service Discovery
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление облачных сервисов
url: /techniques/AML.T0075/
---

После получения доступа злоумышленники могут пытаться перечислить облачные сервисы, запущенные в системе. Эти методы могут различаться в зависимости от модели: platform-as-a-service (PaaS), infrastructure-as-a-service (IaaS), software-as-a-service (SaaS) или AI-as-a-service (AIaaS). У разных облачных провайдеров существует множество сервисов, включая Continuous Integration и Continuous Delivery (CI/CD), Lambda Functions, Entra ID, AI Inference, Generative AI, Agentic AI и другие. Также к ним могут относиться защитные сервисы, такие как AWS GuardDuty и Microsoft Defender for Cloud, и сервисы журналирования, такие как AWS CloudTrail и Google Cloud Audit Logs.

Злоумышленники могут пытаться выявить информацию о сервисах, включенных во всей среде. Инструменты и API Azure, такие как Microsoft Graph API и Azure Resource Manager API, позволяют перечислять ресурсы и сервисы, включая приложения, группы управления, ресурсы и определения политик, а также связи между ними, доступные конкретной учетной записи. Злоумышленники могут использовать инструменты для проверки учетных данных и перечисления ИИ-моделей, доступных в средах разных AIaaS-провайдеров, включая AI21 Labs, Anthropic, AWS Bedrock, Azure, ElevenLabs, MakerSuite, Mistral, OpenAI, OpenRouter и GCP Vertex AI [[sysdig]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0008 Выявление</span><p>Злоумышленники использовали keychecker, чтобы выяснить, какие LLM-сервисы включены в облачной среде и есть ли для этих сервисов квоты ресурсов. Затем злоумышленники проверили, дают ли украденные учетные данные доступ к LLM-ресурсам. Они использовали легитимные запросы `invokeModel` с недопустимым значением -1 для параметра `max_tokens_to_sample`: если учетные данные не давали нужного доступа для вызова модели, такой запрос вызывал ошибку `AccessDenied`. Проверка показала, что украденные учетные данные действительно предоставляли доступ к LLM-ресурсам. Злоумышленники также использовали `GetModelInvocationLoggingConfiguration`, чтобы понять, как настроена модель. Это позволяло им определить, включено ли логирование промптов, и избегать обнаружения при выполнении промптов.</p></a>
</div>


## Источники

- [LLMjacking: Stolen Cloud Credentials Used in New AI Attack | Sysdig](https://www.sysdig.com/blog/llmjacking-stolen-cloud-credentials-used-in-new-ai-attack)
