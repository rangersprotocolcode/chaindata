package log

const (
	DefaultConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/defaultLOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="%Date(2006-01-02 15:04:05.000)[%File:%Line] %Msg%n" />
						</formats>
					</seelog>`

	EVENT = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/event_LOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%Level]|%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`

	MonitorLogConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/monitorLOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)%Msg%n" />
						</formats>
					</seelog>`
	MysqlLogConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/mysqlLOG_INDEX.log" maxsize="300000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`
)
