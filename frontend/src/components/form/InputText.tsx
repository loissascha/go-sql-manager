interface Props {
    id?: string
    value?: string
    onChange?: any
    className?: string
}
export default function InputText({ id, value, onChange, className }: Props) {
    return <input type="text" id={id} value={value} onChange={onChange} className={'bg-gray-700 p-1 rounded w-full ' + className} />
}
