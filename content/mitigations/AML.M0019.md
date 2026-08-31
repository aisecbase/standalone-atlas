---
atlas_id: AML.M0019
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Policy
created_date: "2024-01-12"
description: Требуйте от пользователей подтверждать свою личность перед доступом к продакшен-модели. Требуйте аутентификацию для API-эндпоинтов и отслеживайте запросы к продакшен-модели, чтобы контролировать соблюдение политик...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Control Access to AI Models and Data in Production
technique_count: 20
title: Контроль доступа к ИИ-моделям и данным в продакшене
url: /mitigations/AML.M0019/
---

Требуйте от пользователей подтверждать свою личность перед доступом к продакшен-модели.

Требуйте аутентификацию для API-эндпоинтов и отслеживайте запросы к продакшен-модели, чтобы контролировать соблюдение политик использования и предотвращать злоупотребление моделью.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong><p>Контроль доступа к API моделей может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong><p>Требуйте аутентификации для доступа к ИИ-эндпоинтам в продакшене и отслеживайте запросы, чтобы ограничить зондирование доступных извне ИИ-сервисов без аутентификации.</p></a>
<a class="relation-item" href="/techniques/AML.T0012/"><span class="relation-id">AML.T0012</span><strong>Действующие учетные записи</strong><p>Требуйте аутентификации при доступе к ИИ-эндпоинтам в продакшене и отслеживайте запросы к модели, чтобы выявлять неправомерное использование действительных учётных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0021/"><span class="relation-id">AML.T0021</span><strong>Создание учетных записей</strong><p>Проверяйте идентичность субъекта, прежде чем предоставлять ему доступ к ИИ-системам в продакшене, чтобы вновь созданные учётные записи не получали автоматического доступа к защищённым ИИ-сервисам.</p></a>
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong><p>Злоумышленники могут использовать неограниченный доступ к API, чтобы собрать обучающий набор данных для прокси-модели и раскрыть приватную информацию.</p></a>
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Контроль доступа к API модели может помешать злоумышленнику выполнять чрезмерное количество запросов и выводить систему из строя.</p></a>
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong><p>Контроль доступа может ограничивать доступ к API и предотвращать искусственное увеличение затрат.</p></a>
<a class="relation-item" href="/techniques/AML.T0040/"><span class="relation-id">AML.T0040</span><strong>Доступ к API инференса ИИ-модели</strong><p>Злоумышленники могут использовать неограниченный доступ к API, чтобы получить сведения о продакшен-системе, подготовить атаки и внедрить в систему вредоносные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0042/"><span class="relation-id">AML.T0042</span><strong>Проверка атаки</strong><p>Используйте контроль доступа в продакшене, чтобы помешать злоумышленнику проверять эффективность атаки.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Средства контроля доступа к API модели могут ограничить необходимый злоумышленнику доступ для создания состязательных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><p>Контроль доступа к API модели может лишить злоумышленников доступа, необходимого для методов оптимизации в режиме чёрного ящика.</p></a>
<a class="relation-item" href="/techniques/AML.T0046/"><span class="relation-id">AML.T0046</span><strong>Зашумление ИИ-системы нерелевантными данными</strong><p>Аутентификация для моделей в продакшене может помочь предотвратить анонимный спам шумовыми данными.</p></a>
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong><p>Используйте контроль доступа в продакшене, чтобы помешать злоумышленникам внедрять вредоносные промпты.</p></a>
<a class="relation-item" href="/techniques/AML.T0063/"><span class="relation-id">AML.T0063</span><strong>Выявление выходных данных ИИ-модели</strong><p>Контроль доступа к модели в продакшене может помочь помешать злоумышленникам извлекать информацию из выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong><p>Требуйте аутентификации при доступе к моделям в продакшене и интерфейсам конфигурации и отслеживайте такой доступ.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><p>Требуйте аутентификации при доступе к моделям в продакшене и конфигурации промптов и отслеживайте такой доступ.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><p>Требуйте аутентификации при доступе к моделям в продакшене и конфигурации промптов и отслеживайте такой доступ.</p></a>
<a class="relation-item" href="/techniques/AML.T0091/"><span class="relation-id">AML.T0091</span><strong>Использование альтернативных средств аутентификации</strong><p>Обеспечивайте обязательную авторизацию и отслеживайте использование API ИИ в продакшене на предмет аномальной активности, связанной с повторным предъявлением токенов доступа.</p></a>
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><p>Обеспечивайте обязательную авторизацию и отслеживайте использование API ИИ в продакшене на предмет аномальной активности, связанной с повторным предъявлением токенов доступа.</p></a>
<a class="relation-item" href="/techniques/AML.T0096/"><span class="relation-id">AML.T0096</span><strong>API ИИ-сервиса</strong><p>Аутентифицируйте субъектов, обращающихся к API ИИ-сервиса, и отслеживайте запросы на предмет нарушений политик и случаев неправомерного использования.</p></a>
</div>
