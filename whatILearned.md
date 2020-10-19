## Windows Powershell에서 환경변수 설정
### 환경변수 조회
Get-ChildItem Env:

###환경변수 변경(사용자)
[System.Environment]::SetEnvironmentVariable('FOO', 'bar', [System.EnvironmentVariableTarget]::User)
[System.Environment]::GetEnvironmentVariable('FOO', [System.EnvironmentVariableTarget]::User)

###사용자 환경 변수 삭제
[Environment]::SetEnvironmentVariable('FOO', ' ', 'User')

###환경변수 변경(시스템)
[System.Environment]::SetEnvironmentVariable('FOO', 'bar', System.EnvironmentVariableTarget]::Machine)
[System.Environment]::GetEnvironmentVariable('FOO', [System.EnvironmentVariableTarget]::Machine)
bar

###[Environment]::SetEnvironmentVariable('FOO', 'bar', 'Machine')
[Environment]::GetEnvironmentVariable('FOO', 'Machine')
bar

###[Environment]::SetEnvironmentVariable('FOO', ' ', 'Machine')