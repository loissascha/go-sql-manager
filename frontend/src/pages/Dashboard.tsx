import { useEffect } from 'react'
import { GetDatabaseConfigs } from '../../wailsjs/go/app/App'
import FormGroup from '../components/form/FormGroup'
import FormLabel from '../components/form/FormLabel'
import InputText from '../components/form/InputText'
import './../style.css'

export default function Dashboard() {
    useEffect(() => {
        GetDatabaseConfigs()
            .then((result) => {
                console.log(result)
            })
            .catch((error) => {
                alert(error)
            })
    }, [])

    return (
        <div className="h-full w-full">
            <div className="w-full mb-5">
                <h1 className="text-xl mb-3">Create Connection</h1>
                <div className="grid grid-cols-2 gap-5">
                    <FormGroup>
                        <FormLabel htmlFor="host">Host:</FormLabel>
                        <InputText id="host" />
                    </FormGroup>
                    <FormGroup>
                        <FormLabel>Port:</FormLabel>
                        <InputText />
                    </FormGroup>
                    <FormGroup>
                        <FormLabel>User:</FormLabel>
                        <InputText />
                    </FormGroup>
                    <FormGroup>
                        <FormLabel>Password:</FormLabel>
                        <InputText />
                    </FormGroup>
                </div>
                <div className="mt-5">
                    <FormGroup>
                        <FormLabel>Engine:</FormLabel>
                        <select className="bg-gray-700 text-black w-full">
                            <option>Postgres</option>
                            <option>MySql</option>
                        </select>
                    </FormGroup>
                </div>
            </div>
            <div>Available Connections</div>
        </div>
    )
}
