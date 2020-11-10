window.jsGroup = {}

jsGroup.loop = () => {
    for (let i = 0; i < 10000; i++) { }
}

jsGroup.fibonacci = () => {
    const f = (num) => {
        if (num < 1)
            return 0
        if (num < 3)
            return 1
        return f(num - 1) + f(num - 2)
    }

    f(20)
}

jsGroup.selectionSort = () => {
    const s = (arr) => {
        let temp
        for (let i = 0; i < arr.length; i++) {
            let min = i
            for (j = i + 1; j < arr.length; j++) {
                if (arr[j] < arr[min]) {
                    min = j
                }
            }
            if (min != i) {
                temp = arr[i]
                arr[i] = arr[min]
                arr[min] = temp
            }
        }
    }

    arr = [12, 4124, 3, 24, 5, 43, 34, 654, 7, 54, 23, 643, 34, 76, 32, 125, 43256, 4, 6546, 45, 7, 45, 3, 423, 5, 43, 65, 76, 554, 745, 7]
    s(arr)
}

jsGroup.createElement = () => {
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
    document.createElement('div')
}

jsGroup.appendChild = () => {
    let a = document.createElement('div')
    let b = document.createElement('div')
    let c = document.createElement('div')
    let d = document.createElement('div')
    a.appendChild(b.appendChild(c.appendChild(d)))
}

jsGroup.complex = () => {
    jsGroup.loop()
    jsGroup.fibonacci()
    jsGroup.createElement()
    jsGroup.selectionSort()
    jsGroup.appendChild()
    jsGroup.loop()
    jsGroup.fibonacci()
    jsGroup.createElement()
    jsGroup.selectionSort()
    jsGroup.appendChild()
}