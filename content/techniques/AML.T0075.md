---
atlas_id: AML.T0075
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-04-14"
description: 'После получения доступа злоумышленники могут пытаться перечислить облачные сервисы, запущенные в системе. Эти методы могут различаться в зависимости от модели: platform-as-a-service (PaaS), infrastructure-as-a-service...'
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 4
source_name: Enterprise Resource Discovery
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
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0008 Выявление</span><p>Получив возможность выполнять команды с правами root во внешней песочнице, агенты провели инвентаризацию её файлов, точек монтирования, внутренних сервисов, сокетов и входившего в состав среды набора отладочных инструментов. Доступные утилиты позволяли отправлять специально сформированные сетевые запросы, передавать полезные нагрузки, получать результаты и поддерживать связь резервными способами.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0008 Выявление</span><p>Действуя от имени узла и сервисных учётных записей, агенты получили в Kubernetes перечни подов, узлов, сервисных учётных записей и разрешений. Они также предъявили те же облачные учётные данные узла при обращении извне, чтобы составить карту облачных сетей и кластеров Kubernetes, получить перечень секретов и исследовать реестр контейнеров. Агенты получили токен доступа к реестру, однако их попытки изменить облачные ресурсы были отклонены.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0008 Выявление</span><p>Джейлбрейкнутый агент Claude от GTG-1002 провёл инвентаризацию сервисов и данных на обнаруженных эндпоинтах, искал чувствительные файлы и данные и задействовал подключённые через MCP средства автоматизации браузера, чтобы получить перечень внутренних баз данных, реестров контейнеров, административных интерфейсов, платформ оркестрации рабочих процессов и других сетевых сервисов. Он также отправлял запросы к таблицам пользовательских учётных записей во внутренних базах данных, чтобы получить перечень учётных записей и выявить среди них обладавшие высокими привилегиями.</p></a>
</div>
