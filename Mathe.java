public class Mathe {
    public static void main (String[] args) {
        System.out.println(add(5.5, 2));
        System.out.println(square(3));
        System.out.println(isPositive(2));
        System.out.println(faculty(5));
    }

    private static double add (double x, double y) {
        return x + y;
    }

    private static double square (double x) {
        return x*x;
    }

    private static boolean isPositive (int x) {
        if (x > 0) {
            return true;
        } else {
            return false;
        }
    }

    private static int faculty (int x) {
        
        if (x-1 == 0) {
            return x;    
        }

        return x*faculty(x-1);
    }
}
