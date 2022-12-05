import java.util.*;

public class Day5{
    public static void main(String[] args){
        ArrayList<ArrayDeque<Character>> list = new ArrayList<>();
        Scanner scan = new Scanner(System.in);
        String s = scan.nextLine();
        for (int i = 1; i < s.length(); i+=4) {
            list.add(new ArrayDeque<>());
        }
        while (scan.hasNextLine()) {
            for(int i = 1; i < s.length(); i+=4) {
                if(s.charAt(i) == ' ') {
                    continue;
                }
                else {
                    int size = (i-1)/4;
                    list.get(size).add((s.charAt(i)));
                }
            }
            s = scan.nextLine();
            if(Character.isDigit(s.charAt(1))) break;
        }
        s = scan.nextLine();
        ArrayList<ArrayDeque<Character>> list2 = new ArrayList<>();;
        for (ArrayDeque<Character> stack : list) {
            list2.add(new ArrayDeque<>(stack));
        }

        while (scan.hasNextLine()) {
            String[] strs = scan.nextLine().split(" ");
            int move = Integer.parseInt(strs[1]);
            int from = Integer.parseInt(strs[3])-1;
            int to = Integer.parseInt(strs[5])-1;
            ArrayDeque<Character> temp = new ArrayDeque<>();
            for (int i = 0; i < move; i++) {
                Character c = list.get(from).pop();
                list.get(to).addFirst(c);
                
                Character c2 = list2.get(from).pop();
                list2.get(to).addLast(c2);
                //temp.addFirst(c2);
            }
            // while (!temp.isEmpty()) {
            //     list2.get(to).addFirst(temp.removeFirst());
            // }
        }
        for (ArrayDeque<Character> stack : list) {
            System.out.print(stack.peek());
        }
        System.out.println();
        for (ArrayDeque<Character> stack : list2) {
            System.out.print(stack.peek());
        }
        System.out.println();
    }
}