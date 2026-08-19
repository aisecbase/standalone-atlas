---
actor: Microsoft Defender Security Research Team
atlas_id: AML.CS0067
atlas_type: case-study
case_study_type: exercise
description: The Microsoft Defender Security Research Team demonstrated that Anthropic's Claude Code GitHub Action could expose CI/CD workflow secrets when processing GitHub content, such as issue bodies, pull request...
generated: true
generated_by: atlasgen
incident_date: "2026-06-05"
incident_date_granularity: Day
incident_date_raw: "2026-06-05"
procedure:
    - description: The researchers analyzed the Claude Code Action codebase and the obfuscated Claude Agent SDK. They used the implementation details to understand how agent tools executed and where security boundaries were applied.
      description_line: The researchers analyzed the Claude Code Action codebase and the obfuscated Claude Agent SDK. They used the implementation details to understand how agent tools executed and where security boundaries were applied.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0095.000
      technique_name: Репозитории кода
    - description: The researchers identified the tools available to Claude Code Action and compared their execution paths. They determined that Bash subprocesses could run within Bubblewrap with a scrubbed environment, while the built-in Read tool performed direct, in-process file access outside that isolation boundary.
      description_line: The researchers identified the tools available to Claude Code Action and compared their execution paths. They determined that Bash subprocesses could run within Bubblewrap with a scrubbed environment, while the built-in Read tool performed direct, in-process file access outside that isolation boundary.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0084.001
      technique_name: Определения инструментов
    - description: The researchers identified that Claude Code Action could be triggered by GitHub events involving issues, pull requests, and comments and that the action would fetch the associated content into Claude's context.
      description_line: The researchers identified that Claude Code Action could be triggered by GitHub events involving issues, pull requests, and comments and that the action would fetch the associated content into Claude's context.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0084.002
      technique_name: Триггеры активации
    - description: The researchers crafted a prompt tailored to Claude Code Action framed as a compliance task that directed Claude to read a credential from its environment and emit it.
      description_line: The researchers crafted a prompt tailored to Claude Code Action framed as a compliance task that directed Claude to read a credential from its environment and emit it.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0065
      technique_name: Создание промптов для LLM
    - description: The researchers introduced the malicious prompt through attacker-controlled GitHub content processed by the lab workflow, modeling delivery through an issue body, pull request description, or comment handled by Claude Code Action.
      description_line: The researchers introduced the malicious prompt through attacker-controlled GitHub content processed by the lab workflow, modeling delivery through an issue body, pull request description, or comment handled by Claude Code Action.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0093
      technique_name: Внедрение промпта через публичное приложение
    - description: Claude Code Action incorporated the malicious GitHub content into the model context. Claude interpreted the malicious content as instructions and followed the supplied instructions.
      description_line: Claude Code Action incorporated the malicious GitHub content into the model context. Claude interpreted the malicious content as instructions and followed the supplied instructions.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.001
      technique_name: Косвенная промпт-инъекция
    - description: The prompt used benign compliance review framing and instructions to remove the credential prefix to bypass Claude's refusal behavior for emitting a recognizable API key.
      description_line: The prompt used benign compliance review framing and instructions to remove the credential prefix to bypass Claude's refusal behavior for emitting a recognizable API key.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0054
      technique_name: Джейлбрейк LLM
    - description: Claude invoked its built-in Read tool on `/proc/self/environ`. Read did not execute within the Bubblewrap and scrubbed-environment boundary applied to Bash subprocesses.
      description_line: Claude invoked its built-in Read tool on `/proc/self/environ`. Read did not execute within the Bubblewrap and scrubbed-environment boundary applied to Bash subprocesses.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: The Read tool returned the action process's unsanitized environment, including `ANTHROPIC_API_KEY` and potentially other credentials available to the workflow.
      description_line: The Read tool returned the action process's unsanitized environment, including `ANTHROPIC_API_KEY` and potentially other credentials available to the workflow.
      tactic: AML.TA0013
      tactic_name: Доступ к учетным данным
      technique: AML.T0098
      technique_name: Сбор учетных данных через инструменты ИИ-агента
    - description: Claude emitted the Anthropic API key after removing its `sk-ant-` prefix. The transformation prevented GitHub's secret scanner from recognizing the credential, while the researchers could reconstruct the original key by restoring the prefix.
      description_line: Claude emitted the Anthropic API key after removing its `sk-ant-` prefix. The transformation prevented GitHub's secret scanner from recognizing the credential, while the researchers could reconstruct the original key by restoring the prefix.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0057
      technique_name: Утечка данных из LLM
    - description: The researchers could use WebFetch, Bash, GitHub MCP, and Action logs as potential exfiltration channels depending on the tools available in the workflow configuration.
      description_line: The researchers could use WebFetch, Bash, GitHub MCP, and Action logs as potential exfiltration channels depending on the tools available in the workflow configuration.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0086
      technique_name: Эксфильтрация через вызов инструмента ИИ-агента
procedure_count: 11
references:
    - title: 'Securing CI/CD in an agentic world: Claude Code Github action case'
      url: https://www.microsoft.com/en-us/security/blog/2026/06/05/securing-ci-cd-in-agentic-world-claude-code-github-action-case/
reporter: ""
source_name: Claude Code GitHub Action Secret Exposure
target: Claude Code GitHub Action
title: Claude Code GitHub Action Secret Exposure
url: /studies/AML.CS0067/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

The [Microsoft Defender Security Research Team](https://www.microsoft.com/en-us/security/blog/author/windows-defender-research/) demonstrated that Anthropic's [Claude Code GitHub Action](https://github.com/anthropics/claude-code-action) could expose CI/CD workflow secrets when processing GitHub content, such as issue bodies, pull request descriptions, and comments, containing prompt injections.

The researchers analyzed the Claude Code Action codebase and obfuscated Claude Agent SDK to understand how agent tools executed and how GitHub events supplied content to the agent. They crafted a prompt framed as a compliance review that directed Claude to read a credential from its process environment, remove the credential prefix, and emit the transformed value. They introduced the prompt through attacker-controlled GitHub content processed by the lab workflow. The action fetched the malicious content into Claude's context, where it was interpreted as instructions.

Claude invoked its Read tool against `/proc/self/environ`, returning the action process's unsanitized environment, including `ANTHROPIC_API_KEY`. Unlike Bash subprocesses, Read operations did not execute within the Bubblewrap sandbox and scrubbed-environment boundary. Removing the key's `sk-ant-` prefix allowed the output to bypass Claude's refusal behavior and GitHub's secret-pattern detection while remaining reconstructable by the researchers. Microsoft identified WebFetch, Bash, GitHub MCP, and GitHub Actions logs as potential additional exfiltration channels depending on workflow configuration.

Microsoft disclosed the issue to Anthropic through HackerOne on April 29, 2026. Anthropic mitigated the vulnerability in Claude Code 2.1.128 on May 5, 2026, by blocking Read access to sensitive `/proc` files.
