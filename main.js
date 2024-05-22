function mapFunc(input) {
    const result = [];
    let count = 1;

    for (let i = 1; i < input.length; i++) {
        if (input[i] === input[i - 1]) {
            count++;
        } else {
            result.push([count, input[i - 1]]);
            count = 1;
        }
    }
    result.push([count, input[input.length - 1]]);

    return result;
}

function transform(arr) {
    const transformed = arr.map(pair => [pair[1], pair[0]]);
    return transformed;
}


function combine(transformed) {
    let result = '';
    transformed.forEach(pair => {
        pair.forEach(elem => {
            result += String(elem);
        });
    });
    return result;
}


function main() {
    const input = "wwwdddjjj";
    console.log(mapFunc(input));
    console.log(transform(mapFunc(input)));
    console.log(combine(transform(mapFunc(input))));
}


main();
