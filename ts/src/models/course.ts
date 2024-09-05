export interface Outcome {
    w: string;
    i: string;
    p: string
}

export interface Hole {
    par: number;
    outcomes: { [key: string]: Outcome };
}

export class Course {
    constructor(public name: string, public holes: { [key: string]: Hole }) {}

    printCourseInfo() {
        console.log(`Course Name: ${this.name}`)
        Object.keys(this.holes).forEach(holeNumber => {
            const hole =  this.holes[holeNumber]
            console.log(`Hole ${holeNumber}:  Par ${hole.par}`)
        })
    }

    printHoleOutcomes(holeNumber: string) {
        const hole = this.holes[holeNumber];
        Object.keys(hole.outcomes).forEach(outcomeKey => {
            const outcome = hole.outcomes[outcomeKey];
            console.log(`Outcome ${outcomeKey}: W=${outcome.w}, I=${outcome.i}, P=${outcome.p}`)
        })
    }
}