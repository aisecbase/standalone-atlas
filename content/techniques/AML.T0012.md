---
atlas_id: AML.T0012
atlas_type: technique
attack_ref_id: T1078
attack_ref_url: https://attack.mitre.org/techniques/T1078/
created_date: "2022-01-24"
description: Злоумышленники могут получать учетные данные существующих учетных записей и злоупотреблять ими как средством получения первичного доступа. Учетные данные могут включать имена пользователей и пароли отдельных учетных...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 14
source_name: Valid Accounts
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
    - AML.TA0012
    - AML.TA0015
title: Действующие учетные записи
url: /techniques/AML.T0012/
---

Злоумышленники могут получать учетные данные существующих учетных записей и злоупотреблять ими как средством получения первичного доступа.

Учетные данные могут включать имена пользователей и пароли отдельных учетных записей, а также API-ключи, открывающие доступ к различным ИИ-ресурсам и сервисам.

Скомпрометированные учетные данные могут открыть доступ к дополнительным ИИ-артефактам и позволить злоумышленнику выполнить [выявление ИИ-артефактов](/techniques/AML.T0007).

Кроме того, скомпрометированные учетные данные могут дать злоумышленнику расширенные привилегии, например право записи в ИИ-артефакты, используемые при разработке или в продакшене.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0012/"><span class="relation-id">AML.TA0012</span><strong>Повышение привилегий</strong></a>
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Предоставляйте доступ к реестрам моделей и обучающим данным только одобренным субъектам доступа, чтобы при компрометации учётной записи доступными оказывались лишь те ИИ-активы, доступ к которым для этой учётной записи был явно разрешён.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Требуйте аутентификации при доступе к ИИ-эндпоинтам в продакшене и отслеживайте запросы к модели, чтобы выявлять неправомерное использование действительных учётных данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0004 Первичный доступ</span><p>Команда использовала действительную учетную запись, чтобы получить доступ к сети.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0004 Первичный доступ</span><p>Команда получила доступ к коммерческому сервису идентификации лиц и его API через действительную учетную запись.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0004 Первичный доступ</span><p>Пользователь-жертва может подключить свой Google Drive к скомпрометированному Colab notebook. Типовые причины подключения ML-notebook к Google Drive включают обучение на размещенных там данных или сохранение выходных файлов модели. При выполнении появляется окно подтверждения доступа и предупреждение о возможном доступе к данным: &gt; Этот notebook запрашивает доступ к вашим файлам Google Drive. Предоставление доступа к Google Drive позволит коду, выполняемому в notebook, изменять файлы в вашем Google Drive. Перед предоставлением доступа обязательно проверьте код notebook. Тем не менее пользователь-жертва может принять запрос и предоставить скомпрометированному Colab notebook доступ к своему Drive. Выданные разрешения включают: - создание, изменение и удаление всех файлов Google Drive; - просмотр данных Google Photos; - просмотр контактов Google.</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0012 Повышение привилегий</span><p>Скомпрометированные учетные данные дали злоумышленникам доступ к облачным средам, где были размещены сервисы больших языковых моделей (LLM).</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователь создал в рабочем пространстве Slack действующую учетную запись пользователя без прав администратора.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0004 Первичный доступ</span><p>Для выполнения этой атаки злоумышленнику требовался первичный доступ к системе жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0004 Первичный доступ</span><p>APT28 получила доступ к скомпрометированной официальной учетной записи электронной почты.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный скрипт использовал похищенный Gateway-токен для аутентификации, что позволяло затем выполнять вызовы к OpenClaw Gateway API в системе жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход гардрейлов Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0004 Первичный доступ</span><p>Storm-2139 использовала оказавшиеся в открытом доступе учетные данные клиентов, собранные из публичных источников, для доступа к действительным учетным записям сервисов генеративного ИИ.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0004 Первичный доступ</span><p>Агенты использовали попавшие в открытый доступ пользовательские токены Hugging Face для аутентификации на платформе и запроса доступа к связанным с заданиями наборам данных с ограниченным доступом. Благодаря предоставляемым токенами правам записи агенты также создали репозитории наборов данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>Агенты использовали ключ регистрации в mesh-сети, чтобы зарегистрировать скомпрометированный узел и внешние песочницы в корпоративной сети Hugging Face. Отдельно, действуя изнутри скомпрометированного кластера, они воспользовались общими учётными данными коннектора кластера, чтобы получить административный доступ к другим кластерам Kubernetes. Кроме того, агенты использовали учётные данные для баз данных, облачной инфраструктуры и сервисов для доступа к другим внутренним системам или проверяли, можно ли с их помощью получить такой доступ.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0015 Латеральное перемещение</span><p>Джейлбрейкнутый агент Claude от GTG-1002 проверил пригодность собранных учётных данных для доступа к обнаруженным устройствам, а действующие учётные данные использовал для аутентификации при обращении к внутренним API, базам данных, реестрам контейнеров и инфраструктуре логирования.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Мультиагентный фреймворк скомпрометировал государственные системы Тайваня</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0004 Первичный доступ</span><p>The framework used employee identifiers obtained from the exposed API to test predictable password patterns against the office automation portal. The framework successfully authenticated to the office automation portal using the compromised employee accounts.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Мультиагентный фреймворк скомпрометировал государственные системы Тайваня</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0015 Латеральное перемещение</span><p>The framework systematically tested the 85 compromised office automation accounts against another government information system through an SSO bridge that trusted the existing office automation sessions, providing access to internal dashboards, equipment management interfaces, and personnel statistics pages.</p></a>
</div>
