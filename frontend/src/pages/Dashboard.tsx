import { useEffect, useState } from 'react'
import { ActivateConnection, AddDatabaseConfig, GetDatabaseConfigs } from '../../wailsjs/go/app/App'
import FormGroup from '../components/form/FormGroup'
import FormLabel from '../components/form/FormLabel'
import InputText from '../components/form/InputText'
import Button from '../components/ui/Button'
import Header1 from '../components/ui/Header1'
import './../style.css'

export default function Dashboard() {
    const [createHost, setCreateHost] = useState<string>('')
    const [createPort, setCreatePort] = useState<string>('')
    const [createUser, setCreateUser] = useState<string>('')
    const [createPassword, setCreatePassword] = useState<string>('')
    const [createEngine, setCreateEngine] = useState<string>('')
    const updateCreateHost = (event: any) => setCreateHost(event.target.value)
    const updateCreatePort = (event: any) => setCreatePort(event.target.value)
    const updateCreateUser = (event: any) => setCreateUser(event.target.value)
    const updateCreatePassword = (event: any) => setCreatePassword(event.target.value)
    const updateCreateEngine = (event: any) => setCreateEngine(event.target.value)

    const [connections, setConnections] = useState<any[]>([])

    useEffect(() => {
        loadDbConfigs()
    }, [])

    function loadDbConfigs() {
        GetDatabaseConfigs()
            .then((result) => {
                console.log('result:')
                console.log(result)
                setConnections(result)
            })
            .catch((error) => {
                alert(error)
            })
    }

    function connectionPressed(id: string) {
        ActivateConnection(id)
    }

    function createFormSubmitted(event: any) {
        event.preventDefault()
        AddDatabaseConfig(createHost, createPort, createUser, createPassword, createEngine)
            .then(() => {
                console.log('done!')
                loadDbConfigs()
            })
            .catch((error) => {
                alert(error)
            })
    }

    return (
        <div className="h-full w-full">
            <div className="w-full mb-5">
                <Header1>Create Connection</Header1>
                <form onSubmit={createFormSubmitted}>
                    <div className="grid grid-cols-2 gap-5">
                        <FormGroup>
                            <FormLabel htmlFor="host">Host:</FormLabel>
                            <InputText id="host" value={createHost} onChange={updateCreateHost} />
                        </FormGroup>
                        <FormGroup>
                            <FormLabel htmlFor="port">Port:</FormLabel>
                            <InputText id="port" value={createPort} onChange={updateCreatePort} />
                        </FormGroup>
                        <FormGroup>
                            <FormLabel htmlFor="user">User:</FormLabel>
                            <InputText id="user" value={createUser} onChange={updateCreateUser} />
                        </FormGroup>
                        <FormGroup>
                            <FormLabel htmlFor="password">Password:</FormLabel>
                            <InputText id="password" value={createPassword} onChange={updateCreatePassword} />
                        </FormGroup>
                    </div>
                    <div className="mt-5">
                        <FormGroup>
                            <FormLabel htmlFor="engine">Engine:</FormLabel>
                            <select id="engine" className="bg-gray-700 text-black w-full" value={createEngine} onChange={updateCreateEngine}>
                                <option value={''}>Please select...</option>
                                <option value={'0'}>MySql</option>
                                <option value={'1'}>Postgres</option>
                            </select>
                        </FormGroup>
                    </div>
                    <div className="mt-5">
                        <Button>Save</Button>
                    </div>
                </form>
            </div>
            <Header1>Available Connections</Header1>
            <div>
                {connections.map((connection: any) => (
                    <button
                        key={connection.Id}
                        className="cursor-pointer block w-full text-left py-2 mb-2 bg-gray-600 px-2 rounded"
                        onClick={() => {
                            connectionPressed(connection.Id)
                        }}
                    >
                        <strong>[{connection.Type == 0 ? 'MySql' : connection.Type == 1 ? 'Postgres' : 'Unknown'}]</strong> {connection.Host}:
                        {connection.Port} - {connection.User}
                    </button>
                ))}
            </div>
        </div>
    )
}
