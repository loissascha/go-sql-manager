interface Props {
    children?: any
}
export default function FormGroup({ children }: Props) {
    return <div className="grid grid-cols-[auto_1fr] gap-3 items-center">{children}</div>
}
