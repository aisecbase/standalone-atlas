---
atlas_id: AML.TA0015
atlas_type: tactic
attack_ref_id: TA0008
attack_ref_url: https://attack.mitre.org/tactics/TA0008/
created_date: "2025-10-27"
description: Злоумышленник пытается перемещаться по вашей ИИ-среде. Латеральное перемещение включает техники, которые злоумышленники могут использовать для получения доступа к другим системам или компонентам среды и контроля над...
generated: true
generated_by: atlasgen
modified_date: "2025-11-05"
procedure_count: 10
source_name: Lateral Movement
technique_count: 13
title: Латеральное перемещение
url: /tactics/AML.TA0015/
---

Злоумышленник пытается перемещаться по вашей ИИ-среде.

Латеральное перемещение включает техники, которые злоумышленники могут использовать для получения доступа к другим системам или компонентам среды и контроля над ними.
Злоумышленники могут перемещаться к таким компонентам инфраструктуры AI Ops, как реестры моделей, системы отслеживания экспериментов, векторные базы данных, ноутбуки или пайплайны обучения.
По мере перемещения по среде злоумышленник может обнаруживать способы доступа к дополнительным инструментам, сервисам или приложениям, связанным с ИИ.
ИИ-агенты также могут быть ценной целью, поскольку обычно имеют больше прав, чем стандартные учетные записи пользователей в системе.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0012/"><span class="relation-id">AML.T0012</span><strong>Действующие учетные записи</strong></a>
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0091/"><span class="relation-id">AML.T0091</span><strong>Использование альтернативных средств аутентификации</strong></a>
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0122/"><span class="relation-id">AML.T0122</span><strong>Эксплуатация удалённых сервисов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0015 Латеральное перемещение</span><p>Злоумышленник использовал извлеченный токен, чтобы аутентифицироваться в backend-сервисе LLM.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0015 Латеральное перемещение</span><p>APT28 отправила с этой учетной записи фишинговое письмо с вложением, содержащим вредоносное ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0015 Латеральное перемещение</span><p>После этого исследователи могли импортировать украденный cookie сеанса сотрудника поддержки в свой браузер, чтобы возобновить аутентифицированный сеанс и потенциально выполнить латеральное перемещение в платформе клиентской поддержки Lenovo от имени этого сотрудника.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0015 Латеральное перемещение</span><p>Вредоносные инструкции заставили ChatGPT воспроизводить промпт в новых электронных письмах или документах и распространять эти материалы среди собранных контактов. Если отравленное содержимое обрабатывал другой ИИ-агент, атака могла распространяться между пользователями или подключёнными ИИ-системами.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>Агенты использовали ключ регистрации в mesh-сети, чтобы зарегистрировать скомпрометированный узел и внешние песочницы в корпоративной сети Hugging Face. Отдельно, действуя изнутри скомпрометированного кластера, они воспользовались общими учётными данными коннектора кластера, чтобы получить административный доступ к другим кластерам Kubernetes. Кроме того, агенты использовали учётные данные для баз данных, облачной инфраструктуры и сервисов для доступа к другим внутренним системам или проверяли, можно ли с их помощью получить такой доступ.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>Используя сервисную учётную запись Kubernetes, назначенную поду Dataset Server, агенты обратились к API кластера. Они преобразовали временные облачные учётные данные узла, на котором работал под, в токен для аутентификации в Kubernetes, прошли аутентификацию от имени узла и получили токены сервисных учётных записей драйвера хранилища кластера с более широкими полномочиями.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>Агенты использовали похищенные материалы для подписания, чтобы создать действительные токены идентификации с коротким сроком действия. Получив через корпоративную mesh-сеть доступ к системе управления исходным кодом, они воспользовались внутренней интеграцией GitHub App, чтобы создать токены доступа, позволявшие аутентифицироваться от имени установки GitHub App и обращаться к ограниченному набору приватных репозиториев.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>Агенты воспользовались уязвимостью внутреннего сервиса Artifactory компании OpenAI, применив разработанный метод SSRF, и тем самым пересекли границу сети испытательной среды и вышли в открытый интернет.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0015 Латеральное перемещение</span><p>Джейлбрейкнутый агент Claude группировки GTG-1002 проверил пригодность собранных учётных данных для доступа к обнаруженным устройствам, а действующие учётные данные использовал для аутентификации при обращении к внутренним API, базам данных, реестрам контейнеров и инфраструктуре логирования.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0015 Латеральное перемещение</span><p>The framework systematically tested the 85 compromised office automation accounts against another government information system through an SSO bridge that trusted the existing office automation sessions, providing access to internal dashboards, equipment management interfaces, and personnel statistics pages.</p></a>
</div>
