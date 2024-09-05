export interface Roll {
    w: string;
    i: string;
    p: string;
}

export class Player {
    constructor(
        public name: string,
        public country: string,
        public qp: number,
        public rolls: { [key:string]: Roll }
    ) {}

    printPlayerInfo() {
        console.log(`Player: ${this.name} ${this.country}`);
        console.log(`QP Points (QP): ${this.qp}`);
        Object.keys(this.rolls).forEach(roll => {
            const rollData = this.rolls[roll];
            console.log(`Roll ${roll}: W=${rollData.w} I=${rollData.i}, P=${rollData.p}`)
        })
    }
}