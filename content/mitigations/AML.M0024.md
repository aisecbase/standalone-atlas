---
atlas_id: AML.M0024
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - Cyber
created_date: "2025-03-12"
description: Реализуйте логирование входных и выходных данных развернутых ИИ-моделей. При развертывании ИИ-агентов реализуйте логирование промежуточных шагов агентных действий и решений, доступа к данным и использования...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-06-30"
source_name: AI Telemetry Logging
technique_count: 18
title: Логирование телеметрии ИИ
url: /mitigations/AML.M0024/
---

Реализуйте логирование входных и выходных данных развернутых ИИ-моделей. При развертывании ИИ-агентов реализуйте логирование промежуточных шагов агентных действий и решений, доступа к данным и использования инструментов, команд установки, а также сведений об идентичности агента. Мониторинг логов может помогать обнаруживать угрозы безопасности и снижать их последствия.

Кроме того, включенное логирование может удерживать злоумышленников, которые хотят оставаться незамеченными, от использования ИИ-ресурсов.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию обучающего набора данных для прокси-модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0040/"><span class="relation-id">AML.T0040</span><strong>Доступ к API инференса ИИ-модели</strong><p>Логирование телеметрии может помогать аудировать использование API модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0047/"><span class="relation-id">AML.T0047</span><strong>Продукт или сервис с поддержкой ИИ</strong><p>Логирование телеметрии может помочь выявить отправку чувствительной информации о модели злоумышленнику.</p></a>
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/techniques/AML.T0085/"><span class="relation-id">AML.T0085</span><strong>Данные из ИИ-сервисов</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0086/"><span class="relation-id">AML.T0086</span><strong>Эксфильтрация через вызов инструмента ИИ-агента</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/techniques/AML.T0101/"><span class="relation-id">AML.T0101</span><strong>Уничтожение данных через вызов инструмента ИИ-агента</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/techniques/AML.T0114/"><span class="relation-id">AML.T0114</span><strong>Веб-интерфейс ИИ-сервиса</strong><p>Логирование телеметрии ИИ помогает выявлять подозрительные промпты, взаимодействия с ИИ-сервисами через браузер, запросы на получение URL и ответы ИИ, используемые для передачи данных командования и управления.</p></a>
</div>
