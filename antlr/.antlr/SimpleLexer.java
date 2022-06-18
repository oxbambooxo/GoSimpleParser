// Generated from /Users/db.huang/Resources/GoSimpleScript2/SimpleLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SimpleLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Int=1, IntegerLiteral=2, AssignmentOP=3, Plus=4, Minus=5, Star=6, Slash=7, 
		SemiColon=8, LeftParen=9, RightParen=10, Identifier=11, Whitespace=12, 
		Newline=13;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star", "Slash", 
			"SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace", "Newline"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", null, "'='", "'+'", "'-'", "'*'", "'/'", "';'", "'('", 
			"')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star", 
			"Slash", "SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace", 
			"Newline"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public SimpleLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SimpleLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\17M\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\2\3\3\6\3#\n\3\r\3\16\3$\3"+
		"\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3"+
		"\f\7\f9\n\f\f\f\16\f<\13\f\3\r\6\r?\n\r\r\r\16\r@\3\r\3\r\3\16\3\16\5"+
		"\16G\n\16\3\16\5\16J\n\16\3\16\3\16\2\2\17\3\3\5\4\7\5\t\6\13\7\r\b\17"+
		"\t\21\n\23\13\25\f\27\r\31\16\33\17\3\2\6\3\2\62;\5\2C\\aac|\6\2\62;C"+
		"\\aac|\4\2\13\13\"\"\2Q\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2"+
		"\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\3\35\3\2\2\2\5\"\3\2\2"+
		"\2\7&\3\2\2\2\t(\3\2\2\2\13*\3\2\2\2\r,\3\2\2\2\17.\3\2\2\2\21\60\3\2"+
		"\2\2\23\62\3\2\2\2\25\64\3\2\2\2\27\66\3\2\2\2\31>\3\2\2\2\33I\3\2\2\2"+
		"\35\36\7k\2\2\36\37\7p\2\2\37 \7v\2\2 \4\3\2\2\2!#\t\2\2\2\"!\3\2\2\2"+
		"#$\3\2\2\2$\"\3\2\2\2$%\3\2\2\2%\6\3\2\2\2&\'\7?\2\2\'\b\3\2\2\2()\7-"+
		"\2\2)\n\3\2\2\2*+\7/\2\2+\f\3\2\2\2,-\7,\2\2-\16\3\2\2\2./\7\61\2\2/\20"+
		"\3\2\2\2\60\61\7=\2\2\61\22\3\2\2\2\62\63\7*\2\2\63\24\3\2\2\2\64\65\7"+
		"+\2\2\65\26\3\2\2\2\66:\t\3\2\2\679\t\4\2\28\67\3\2\2\29<\3\2\2\2:8\3"+
		"\2\2\2:;\3\2\2\2;\30\3\2\2\2<:\3\2\2\2=?\t\5\2\2>=\3\2\2\2?@\3\2\2\2@"+
		">\3\2\2\2@A\3\2\2\2AB\3\2\2\2BC\b\r\2\2C\32\3\2\2\2DF\7\17\2\2EG\7\f\2"+
		"\2FE\3\2\2\2FG\3\2\2\2GJ\3\2\2\2HJ\7\f\2\2ID\3\2\2\2IH\3\2\2\2JK\3\2\2"+
		"\2KL\b\16\2\2L\34\3\2\2\2\t\2$8:@FI\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}