interface Props {
    title: string
    onClick?: any
    selected?: boolean
}
export default function DbListItem({ title, onClick, selected }: Props) {
    return (
        <div
            className={'w-full mx-auto p-3 text-center cursor-pointer rounded ' + (selected ? 'bg-gray-400' : 'bg-gray-500')}
            onClick={() => {
                if (onClick) {
                    onClick()
                }
            }}
        >
            {title}
        </div>
    )
}
