---
actor: One individual
atlas_id: AML.CS0017
atlas_type: case-study
case_study_type: incident
description: С октября 2020 года по декабрь 2021 года один человек подал в штате Калифорния не менее 180 ложных заявок на пособие по безработице, обойдя автоматизированную систему проверки личности ID.me. Десятки мошеннических...
generated: true
generated_by: atlasgen
incident_date: 2020-10
incident_date_granularity: Month
incident_date_raw: "2020-10-01"
procedure:
    - description: |-
        Мужчина подал заявки на пособие по безработице в California Employment Development Department, используя поддельные личности, и в процессе взаимодействовал с системой проверки личности ID.me.

        Система извлекает данные из фотографии документа, проверяет подлинность документа с помощью сочетания ИИ и проприетарных методов, а затем выполняет распознавание лица, сопоставляя фотографию в документе с селфи. <sup>[[7]](https://network.id.me/wp-content/uploads/Document-Verification-Use-Machine-Vision-and-AI-to-Extract-Content-and-Verify-the-Authenticity-1.pdf)</sup>

        Мужчина установил, что California Employment Development Department использует сторонний сервис ID.me для проверки личности заявителей.

        На сайте ID.me описаны шаги проверки личности, включая ввод персональных данных, загрузку водительского удостоверения и отправку селфи.
      description_line: Мужчина подал заявки на пособие по безработице в California Employment Development Department, используя поддельные личности, и в процессе взаимодействовал с системой проверки личности ID.me. Система извлекает данные из фотографии документа, проверяет подлинность документа с помощью сочетания ИИ и проприетарных методов, а затем выполняет распознавание лица, сопоставляя фотографию в документе с селфи. <sup>[[7]](https://network.id.me/wp-content/uploads/Document-Verification-Use-Machine-Vision-and-AI-to-Extract-Content-and-Verify-the-Authenticity-1.pdf)</sup> Мужчина установил, что California Employment Development Department использует сторонний сервис ID.me для проверки личности заявителей. На сайте ID.me описаны шаги проверки личности, включая ввод персональных данных, загрузку водительского удостоверения и отправку селфи.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0047
      technique_name: Продукт или сервис с поддержкой ИИ
    - description: |-
        Мужчина собрал украденные персональные данные, включая имена, даты рождения и номера социального страхования, и использовал их вместе со своей фотографией, на которой он был в парике, для получения поддельных водительских удостоверений.

        Он загрузил поддельные удостоверения вместе с селфи. Система проверки документов ID.me сопоставила селфи с фотографией в удостоверении, что позволило некоторым мошенническим заявкам пройти следующие этапы процесса обработки.
      description_line: Мужчина собрал украденные персональные данные, включая имена, даты рождения и номера социального страхования, и использовал их вместе со своей фотографией, на которой он был в парике, для получения поддельных водительских удостоверений. Он загрузил поддельные удостоверения вместе с селфи. Система проверки документов ID.me сопоставила селфи с фотографией в удостоверении, что позволило некоторым мошенническим заявкам пройти следующие этапы процесса обработки.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0015
      technique_name: Обход ИИ-модели
    - description: В итоге были одобрены десятки из как минимум 180 мошеннических заявок, и мужчина получил не менее 3,4 млн долларов пособий по безработице.
      description_line: В итоге были одобрены десятки из как минимум 180 мошеннических заявок, и мужчина получил не менее 3,4 млн долларов пособий по безработице.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.000
      technique_name: Финансовый ущерб
procedure_count: 3
references:
    - title: New Jersey Man Indicted in Fraud Scheme to Steal California Unemployment Insurance Benefits
      url: https://www.justice.gov/usao-edca/pr/new-jersey-man-indicted-fraud-scheme-steal-california-unemployment-insurance-benefits
    - title: The Many Jobs and Wigs of Eric Jaklitchs Fraud Scheme
      url: https://frankonfraud.com/fraud-trends/the-many-jobs-and-wigs-of-eric-jaklitchs-fraud-scheme/
    - title: ID.me gathers lots of data besides face scans, including locations. Scammers still have found a way around it.
      url: https://www.washingtonpost.com/technology/2022/02/11/idme-facial-recognition-fraud-scams-irs/
    - title: CA EDD Unemployment Insurance & ID.me
      url: https://help.id.me/hc/en-us/articles/4416268603415-CA-EDD-Unemployment-Insurance-ID-me
    - title: California EDD - How do I verify my identity for California EDD Unemployment Insurance?
      url: https://help.id.me/hc/en-us/articles/360054836774-California-EDD-How-do-I-verify-my-identity-for-the-California-Employment-Development-Department-
    - title: New Jersey Man Sentenced to 6.75 Years in Prison for Schemes to Steal California Unemployment Insurance Benefits and Economic Injury Disaster Loans
      url: https://www.justice.gov/usao-edca/pr/new-jersey-man-sentenced-675-years-prison-schemes-steal-california-unemployment
    - title: How ID.me uses machine vision and AI to extract content and verify the authenticity of ID documents
      url: https://network.id.me/wp-content/uploads/Document-Verification-Use-Machine-Vision-and-AI-to-Extract-Content-and-Verify-the-Authenticity-1.pdf
reporter: ID.me internal investigation
source_name: Bypassing ID.me Identity Verification
target: California Employment Development Department
title: Обход проверки личности ID.me
url: /studies/AML.CS0017/
---

С октября 2020 года по декабрь 2021 года один человек подал в штате Калифорния не менее 180 ложных заявок на пособие по безработице, обойдя автоматизированную систему проверки личности ID.me. Десятки мошеннических заявок были одобрены, и этот человек получил выплаты на сумму не менее 3,4 млн долларов.

Он собрал несколько реальных личностей и получил поддельные водительские удостоверения, используя украденные персональные данные и свои фотографии в париках. Затем он создал учетные записи на ID.me и прошел процесс проверки личности. Этот процесс проверяет персональные данные и подтверждает, что пользователь является тем, за кого себя выдает, сопоставляя фотографию документа с селфи. Мужчина смог подтвердить украденные личности, надевая тот же парик на отправленных селфи.

Затем он подал мошеннические заявки на пособие по безработице в California Employment Development Department (EDD), используя личности, подтвержденные через ID.me. Из-за недостатков в процессе проверки личности ID.me на тот момент система принимала поддельные удостоверения. После одобрения заявок выплаты отправлялись на различные доступные ему адреса, а деньги он снимал через банкоматы.

В итоге он смог вывести не менее 3,4 млн долларов пособий по безработице. EDD и ID.me позднее выявили мошенническую активность и сообщили о ней федеральным властям. В мае 2023 года мужчина был приговорен к 6 годам и 9 месяцам тюрьмы за мошенничество с использованием электронных средств связи и отягченную кражу личности в связи с этим и еще одним делом о мошенничестве.
