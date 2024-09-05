import { Course } from '../models/course'

describe('Course details', () => {
    it('should have course attributes', () => {
        const course = new Course('Augusta National', {});
        expect(course.name).toBe('Augusta National')
    })

    // add a test to check if the course has holes
    it('should have holes', () => {
        const course = new Course('Augusta National', {
            '1': {
                par: 4,
                outcomes: {
                    '1': { w: '1', i: '2', p: '3' },
                    '2': { w: '2', i: '3', p: '4' }
                }
            }
        });

        expect(Object.keys(course.holes).length).toBe(1);

        // check if the hole has the correct par
        expect(course.holes['1'].par).toBe(4);

        // check that there are two outcomes
        expect(Object.keys(course.holes['1'].outcomes).length).toBe(2); 
    })
})