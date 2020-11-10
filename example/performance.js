function main() {
    testGroup('loop', [
        js(jsGroup.loop),
        wasm(wasmGroup.loop)
    ])

    testGroup('fibonacci', [
        js(jsGroup.fibonacci),
        wasm(wasmGroup.fibonacci)
    ])

    testGroup('selection sort', [
        js(jsGroup.selectionSort),
        wasm(wasmGroup.selectionSort)
    ])

    testGroup('create element', [
        js(jsGroup.createElement),
        wasm(wasmGroup.createElement)
    ])

    testGroup('append child', [
        js(jsGroup.appendChild),
        wasm(wasmGroup.appendChild)
    ])

    testGroup('complex', [
        js(jsGroup.complex),
        wasm(wasmGroup.complex)
    ])
}

function js(func) {
    return {
        func,
        name: 'javascript'
    }
}

function wasm(func) {
    return {
        func,
        name: 'webAssembly'
    }
}

function testGroup(groupName, arr) {
    if (!(arr instanceof Array)) return
    console.log(`%c------${groupName}------`, "color:red")
    console.log(`%csync:`, "color:blue")
    arr.forEach((item) => {
        testSync(item.name, item.func)
    })
    console.log(`%c------test end------`, "color:green")
}

async function testAsyncGroup(groupName, arr) {
    if (!(arr instanceof Array)) return
    console.log(`%c------${groupName}------`, "color:red")
    console.log(`%casync:`, "color:blue")
    arr.forEach(async (item) => {
        await testAsync(item.name, item.func)
    })
    console.log(`%c------test end------`, "color:green")
}

function testSync(name, func) {
    console.time(name)
    func()
    console.timeEnd(name)
}

async function testAsync(name, funcAsync) {
    console.time(name)
    await funcAsync()
    console.timeEnd(name)
}

function initializeWebAssembly(path) {
    if (!WebAssembly.instantiateStreaming) {
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer()
            return await WebAssembly.instantiate(source, importObject)
        }
    }

    const go = new Go()

    return new Promise((resolve, reject) => {
        WebAssembly.instantiateStreaming(fetch(path), go.importObject)
            .then((res) => {
                console.timeEnd("initialize")
                go.run(res.instance)
            })
            .then(() => {
                resolve()
            })
            .catch(() => {
                reject()
            })
    })
}

console.time("initialize")
initializeWebAssembly("main.wasm")