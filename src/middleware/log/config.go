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

	P2PRawConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/p2prawLOG_INDEX.log" maxsize="1000000000" maxrolls="10"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`

	P2PLogConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/p2pLOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`

	P2PBizLogConfig = `<seelog minlevel="info">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/p2p_bizLOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`
	MiddlewareLogConfig = `<seelog minlevel="error">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/middlewareLOG_INDEX.log" maxsize="100000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`

	PerformanceLogConfig = `<seelog minlevel="info">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/performanceLOG_INDEX.log" maxsize="200000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
						</formats>
					</seelog>`

	LdbLogConfig = `<seelog minlevel="debug">
						<outputs formatid="default">
							<rollingfile type="size" filename="./logs/ldbLOG_INDEX.log" maxsize="50000000" maxrolls="1"/>
						</outputs>
						<formats>
							<format id="default" format="[%LEV]%Date(2006-01-02 15:04:05.000)[%File:%Line]%Msg%n" />
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
