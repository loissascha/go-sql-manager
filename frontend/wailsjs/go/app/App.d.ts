// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {configs} from '../models';

export function ActivateConnection(arg1:string):Promise<void>;

export function AddDatabaseConfig(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<void>;

export function GetDatabaseConfigs():Promise<Array<configs.DatabaseConfig>>;

export function ListDbTables():Promise<Array<string>>;

export function ListTable(arg1:string,arg2:string):Promise<string>;

export function ListTables(arg1:string):Promise<Array<string>>;
