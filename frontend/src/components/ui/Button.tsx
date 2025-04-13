interface Props {
    children?: any
    onClick?: any
}
export default function Button({ children, onClick }: Props) {
    return (
        <button
            onClick={() => {
                if (onClick) onClick()
            }}
            className="cursor-pointer bg-blue-500 hover:bg-blue-600 px-3 py-2 rounded"
        >
            {children}
        </button>
    )
}
