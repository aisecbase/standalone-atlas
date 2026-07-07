---
actor: APT28
atlas_id: AML.CS0044
atlas_type: case-study
case_study_type: ""
description: В июле 2025 года украинские органы сообщили о появлении LAMEHUG, нового вредоносного ПО на базе ИИ, которое они атрибутировали APT28, актору угроз, по их оценке связанному с российским государством и также...
generated: true
generated_by: atlasgen
incident_date: ""
incident_date_granularity: ""
incident_date_raw: ""
procedure:
    - description: APT28 получила доступ к скомпрометированной официальной учетной записи электронной почты.
      description_line: APT28 получила доступ к скомпрометированной официальной учетной записи электронной почты.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: APT28 отправила с этой учетной записи фишинговое письмо с вложением, содержащим вредоносное ПО.
      description_line: APT28 отправила с этой учетной записи фишинговое письмо с вложением, содержащим вредоносное ПО.
      tactic: AML.TA0015
      tactic_name: Латеральное перемещение
      technique: AML.T0052
      technique_name: Фишинг
    - description: В письме злоумышленники выдавали себя за представителя государственного министерства.
      description_line: В письме злоумышленники выдавали себя за представителя государственного министерства.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0073
      technique_name: Имперсонация
    - description: Вложение называлось `Appendix.pdf.zip`, что могло заставить получателя принять его за легитимный PDF-файл.
      description_line: Вложение называлось `Appendix.pdf.zip`, что могло заставить получателя принять его за легитимный PDF-файл.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0074
      technique_name: Маскировка
    - description: Вложение содержало исполняемый файл с расширением `.pif`, созданный из Python-кода с помощью PyInstaller. CERT-UA классифицировал его как LAMEHUG. Файлы с расширением `.pif` являются исполняемыми в Windows.
      description_line: Вложение содержало исполняемый файл с расширением `.pif`, созданный из Python-кода с помощью PyInstaller. CERT-UA классифицировал его как LAMEHUG. Файлы с расширением `.pif` являются исполняемыми в Windows.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0011
      technique_name: Запуск пользователем
    - description: LAMEHUG злоупотреблял Hugging Face API модели Qwen 2.5 Coder 32B Instruct, чтобы генерировать вредоносные команды из промптов на естественном языке.
      description_line: LAMEHUG злоупотреблял Hugging Face API модели Qwen 2.5 Coder 32B Instruct, чтобы генерировать вредоносные команды из промптов на естественном языке.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0102
      technique_name: Генерация вредоносных команд
    - description: LAMEHUG использовал команды, сгенерированные ИИ, для сбора сведений о системе с сохранением в `%PROGRAMDATA%\info\info.txt`, а также рекурсивно просматривал папки Documents, Desktop и Downloads, чтобы подготовить файлы к эксфильтрации.
      description_line: LAMEHUG использовал команды, сгенерированные ИИ, для сбора сведений о системе с сохранением в `%PROGRAMDATA%\info\info.txt`, а также рекурсивно просматривал папки Documents, Desktop и Downloads, чтобы подготовить файлы к эксфильтрации.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0037
      technique_name: Данные из локальной системы
    - description: LAMEHUG эксфильтровал собранные данные на серверы под контролем злоумышленников через SFTP или HTTP POST-запросы.
      description_line: LAMEHUG эксфильтровал собранные данные на серверы под контролем злоумышленников через SFTP или HTTP POST-запросы.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
procedure_count: 8
references:
    - title: LameHug malware uses AI LLM to craft Windows data-theft commands in real-time
      url: https://www.bleepingcomputer.com/news/security/lamehug-malware-uses-ai-llm-to-craft-windows-data-theft-commands-in-real-time/
    - title: UAC-0001 cyberattacks on the security and defense sector using the LAMEHUG software tool, which uses LLM (large language model) (CERT-UA#16039)
      url: https://cert.gov.ua/article/6284730
    - title: 'APT28''s New Arsenal: LAMEHUG, the First AI-Powered Malware'
      url: https://logpoint.com/en/blog/apt28s-new-arsenal-lamehug-the-first-ai-powered-malware
reporter: CERT-UA
source_name: 'LAMEHUG: Malware Leveraging Dynamic AI-Generated Commands'
target: Ukraine's security and defense sector
title: 'LAMEHUG: вредоносное ПО, использующее динамические команды, сгенерированные ИИ'
url: /studies/AML.CS0044/
---

В июле 2025 года украинские органы сообщили о появлении LAMEHUG, нового вредоносного ПО на базе ИИ, которое они атрибутировали [APT28](https://attack.mitre.org/groups/G0007/), актору угроз, по их оценке связанному с российским государством и также отслеживаемому как Forest Blizzard или UAC-0001. LAMEHUG использует большую языковую модель (LLM), чтобы динамически генерировать команды на зараженных хостах.

Кампания началась с фишинговой атаки: злоумышленники использовали скомпрометированную государственную учетную запись электронной почты, чтобы доставить вредоносный ZIP-архив, замаскированный под `Appendix.pdf.zip`. Архив содержал вредоносное ПО LAMEHUG, исполняемый файл на базе Python, упакованный с помощью PyInstaller. При выполнении вредоносное ПО обращается к эндпоинту LLM, чтобы генерировать вредоносные команды из промптов на естественном языке. Динамически сгенерированные команды могут усложнять обнаружение вредоносного ПО. LAMEHUG был настроен на сбор файлов из локальной системы и их эксфильтрацию.
