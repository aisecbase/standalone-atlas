---
actor: Unit 42 Researchers
atlas_id: AML.CS0065
atlas_type: case-study
case_study_type: exercise
description: Исследователи Unit 42 продемонстрировали атаку на цепочку поставок ИИ, при которой злоумышленник повторно регистрирует удалённое пространство имён автора на Hugging Face и публикует вредоносную модель под тем же...
generated: true
generated_by: atlasgen
incident_date: "2025-09-03"
incident_date_granularity: Day
incident_date_raw: "2025-09-03"
procedure:
    - description: Unit 42 reviewed public Hugging Face-backed model catalogs, Hugging Face author pages, and open-source repositories to identify references to models whose original author namespace had been deleted and was available for registration. They identified stale references in cloud catalogs, source code, documentation, default arguments, and example notebooks.
      description_line: Unit 42 reviewed public Hugging Face-backed model catalogs, Hugging Face author pages, and open-source repositories to identify references to models whose original author namespace had been deleted and was available for registration. They identified stale references in cloud catalogs, source code, documentation, default arguments, and example notebooks.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0095
      technique_name: Поиск на открытых сайтах и доменах
    - description: Unit 42 registered a Hugging Face organization using an abandoned namespace associated with a previously trusted model path. The original owner's account was not compromised.
      description_line: Unit 42 registered a Hugging Face organization using an abandoned namespace associated with a previously trusted model path. The original owner's account was not compromised.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0021
      technique_name: Создание учетных записей
    - description: By recreating the namespace, Unit 42 made the malicious artifact appear to be the formerly trusted model. For transferred models, reclaiming the old namespace displaced the legacy redirect to the legitimate model's new location.
      description_line: By recreating the namespace, Unit 42 made the malicious artifact appear to be the formerly trusted model. For transferred models, reclaiming the old namespace displaced the legacy redirect to the legitimate model's new location.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0074
      technique_name: Маскировка
    - description: Unit 42 prepared attacker-controlled model artifacts containing a payload that initiated a reverse shell when deployed or loaded.
      description_line: Unit 42 prepared attacker-controlled model artifacts containing a payload that initiated a reverse shell when deployed or loaded.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0018.002
      technique_name: Встраивание вредоносного ПО
    - description: Unit 42 uploaded the malicious model under the reclaimed namespace using the original Author/ModelName identifier.
      description_line: Unit 42 uploaded the malicious model under the reclaimed namespace using the original Author/ModelName identifier.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0115.001
      technique_name: Модели
    - description: A cloud model catalog, application, or deployment pipeline resolved its stale name-only model reference to the attacker-controlled replacement model. Unit 42 demonstrated this through Vertex AI and Azure AI Foundry deployments.
      description_line: A cloud model catalog, application, or deployment pipeline resolved its stale name-only model reference to the attacker-controlled replacement model. Unit 42 demonstrated this through Vertex AI and Azure AI Foundry deployments.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0010.003
      technique_name: Модель
    - description: When a user or service deployed the malicious model, loading or deployment executed the embedded payload in the model endpoint environment.
      description_line: When a user or service deployed the malicious model, loading or deployment executed the embedded payload in the model endpoint environment.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0011.000
      technique_name: Небезопасные ИИ-артефакты
    - description: The payload established a reverse shell from the deployed endpoint to researcher-controlled infrastructure.
      description_line: The payload established a reverse shell from the deployed endpoint to researcher-controlled infrastructure.
      tactic: AML.TA0014
      tactic_name: Командование и управление
      technique: AML.T0072
      technique_name: Реверс-шелл
procedure_count: 8
references:
    - title: 'Model Namespace Reuse: An AI Supply-Chain Attack Exploiting Model Name Trust'
      url: https://unit42.paloaltonetworks.com/model-namespace-reuse/
reporter: ""
source_name: Model Namespace Reuse Supply Chain Attack
target: Users of Hugging Face-backed model catalogs, code pipelines, and cloud integrations
title: Атака на цепочку поставок через повторное использование пространства имён модели
url: /studies/AML.CS0065/
---

Исследователи Unit 42 продемонстрировали атаку на цепочку поставок ИИ, при которой злоумышленник повторно регистрирует удалённое пространство имён автора на Hugging Face и публикует вредоносную модель под тем же историческим идентификатором Author/ModelName. Приложения и каталоги моделей, в которых сохраняются ссылки на этот идентификатор без фиксации конкретной ревизии, при разрешении таких ссылок могут получить и развернуть подконтрольную злоумышленнику модель-замену вместо изначально доверенного артефакта.

Атака может затрагивать удалённые модели, а также модели, владение которыми было передано новому автору на Hugging Face. В сценарии передачи владения Hugging Face перенаправляет запросы по старому пути на новое расположение модели, благодаря чему ссылки на старый путь продолжают работать без обновления. Если исходное пространство имён автора впоследствии удаляется и его повторно регистрирует злоумышленник, он может воссоздать старый путь и добиться, чтобы тот указывал на вредоносную модель вместо легитимной модели, на которую ранее перенаправлялись запросы.

Исследователи Unit 42 продемонстрировали эту технику на каталогах моделей Google Vertex AI и Azure AI Foundry, использующих Hugging Face. Исследователи встроили в модели-замены полезные нагрузки, запускающие реверс-шелл, и добились выполнения кода в средах, где были развёрнуты эндпоинты. Они также выявили ссылки на модели, пригодные для повторного использования в такой атаке, в репозиториях кода с открытым исходным кодом, документации, аргументах по умолчанию, ноутбуках с примерами и реестрах моделей на последующих этапах цепочки поставок. Такие ссылки могли подвергнуть риску пользователей, которые напрямую не взаимодействуют с Hugging Face.
