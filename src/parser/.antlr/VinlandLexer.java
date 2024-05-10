// Generated from c:/Users/sunee/Code/vinland/src/parser/Vinland.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class VinlandLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, Whitespace=12, Identifier=13, IntLiteral=14, DecimalNumber=15, 
		HexadecimalNumber=16, BinaryNumber=17, FloatLiteral=18, StringLiteral=19, 
		CharEscapeSeq=20, BooleanLiteral=21, True=22, False=23;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "Whitespace", "Identifier", "IntLiteral", "DecimalNumber", 
			"HexadecimalNumber", "BinaryNumber", "FloatLiteral", "StringLiteral", 
			"StringElement", "CharEscapeSeq", "BooleanLiteral", "True", "False"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'='", "'fn'", "'('", "')'", "','", "'+'", "'-'", 
			"'*'", "'/'", null, null, null, null, null, null, null, null, null, null, 
			"'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"Whitespace", "Identifier", "IntLiteral", "DecimalNumber", "HexadecimalNumber", 
			"BinaryNumber", "FloatLiteral", "StringLiteral", "CharEscapeSeq", "BooleanLiteral", 
			"True", "False"
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


	public VinlandLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Vinland.g4"; }

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
		"\u0004\u0000\u0017\u009e\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001"+
		"\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0004\u000bJ\b\u000b\u000b"+
		"\u000b\f\u000bK\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0005\fR\b\f\n"+
		"\f\f\fU\t\f\u0001\r\u0001\r\u0001\r\u0003\rZ\b\r\u0001\u000e\u0003\u000e"+
		"]\b\u000e\u0001\u000e\u0004\u000e`\b\u000e\u000b\u000e\f\u000ea\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0004\u000fh\b\u000f\u000b"+
		"\u000f\f\u000fi\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0004"+
		"\u0010p\b\u0010\u000b\u0010\f\u0010q\u0001\u0011\u0004\u0011u\b\u0011"+
		"\u000b\u0011\f\u0011v\u0001\u0011\u0001\u0011\u0005\u0011{\b\u0011\n\u0011"+
		"\f\u0011~\t\u0011\u0001\u0012\u0001\u0012\u0005\u0012\u0082\b\u0012\n"+
		"\u0012\f\u0012\u0085\t\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u008b\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0015\u0001\u0015\u0003\u0015\u0092\b\u0015\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0000\u0000\u0018\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\'\u0000)\u0014+\u0015-\u0016/\u0017\u0001\u0000\b\u0003"+
		"\u0000\t\n\r\r  \u0002\u0000AZaz\u0003\u000009AZaz\u0001\u000009\u0002"+
		"\u000009AF\u0001\u000001\u0002\u0000 !#\u007f\u0005\u0000\"\"\\\\nnrr"+
		"tt\u00a9\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000"+
		"\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000"+
		"\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000"+
		"\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000"+
		"\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000"+
		"\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000"+
		"\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000"+
		"\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000"+
		"!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001"+
		"\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000\u0000"+
		"\u0000\u0000-\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000\u0001"+
		"1\u0001\u0000\u0000\u0000\u00033\u0001\u0000\u0000\u0000\u00055\u0001"+
		"\u0000\u0000\u0000\u00077\u0001\u0000\u0000\u0000\t:\u0001\u0000\u0000"+
		"\u0000\u000b<\u0001\u0000\u0000\u0000\r>\u0001\u0000\u0000\u0000\u000f"+
		"@\u0001\u0000\u0000\u0000\u0011B\u0001\u0000\u0000\u0000\u0013D\u0001"+
		"\u0000\u0000\u0000\u0015F\u0001\u0000\u0000\u0000\u0017I\u0001\u0000\u0000"+
		"\u0000\u0019O\u0001\u0000\u0000\u0000\u001bY\u0001\u0000\u0000\u0000\u001d"+
		"\\\u0001\u0000\u0000\u0000\u001fc\u0001\u0000\u0000\u0000!k\u0001\u0000"+
		"\u0000\u0000#t\u0001\u0000\u0000\u0000%\u007f\u0001\u0000\u0000\u0000"+
		"\'\u008a\u0001\u0000\u0000\u0000)\u008c\u0001\u0000\u0000\u0000+\u0091"+
		"\u0001\u0000\u0000\u0000-\u0093\u0001\u0000\u0000\u0000/\u0098\u0001\u0000"+
		"\u0000\u000012\u0005{\u0000\u00002\u0002\u0001\u0000\u0000\u000034\u0005"+
		"}\u0000\u00004\u0004\u0001\u0000\u0000\u000056\u0005=\u0000\u00006\u0006"+
		"\u0001\u0000\u0000\u000078\u0005f\u0000\u000089\u0005n\u0000\u00009\b"+
		"\u0001\u0000\u0000\u0000:;\u0005(\u0000\u0000;\n\u0001\u0000\u0000\u0000"+
		"<=\u0005)\u0000\u0000=\f\u0001\u0000\u0000\u0000>?\u0005,\u0000\u0000"+
		"?\u000e\u0001\u0000\u0000\u0000@A\u0005+\u0000\u0000A\u0010\u0001\u0000"+
		"\u0000\u0000BC\u0005-\u0000\u0000C\u0012\u0001\u0000\u0000\u0000DE\u0005"+
		"*\u0000\u0000E\u0014\u0001\u0000\u0000\u0000FG\u0005/\u0000\u0000G\u0016"+
		"\u0001\u0000\u0000\u0000HJ\u0007\u0000\u0000\u0000IH\u0001\u0000\u0000"+
		"\u0000JK\u0001\u0000\u0000\u0000KI\u0001\u0000\u0000\u0000KL\u0001\u0000"+
		"\u0000\u0000LM\u0001\u0000\u0000\u0000MN\u0006\u000b\u0000\u0000N\u0018"+
		"\u0001\u0000\u0000\u0000OS\u0007\u0001\u0000\u0000PR\u0007\u0002\u0000"+
		"\u0000QP\u0001\u0000\u0000\u0000RU\u0001\u0000\u0000\u0000SQ\u0001\u0000"+
		"\u0000\u0000ST\u0001\u0000\u0000\u0000T\u001a\u0001\u0000\u0000\u0000"+
		"US\u0001\u0000\u0000\u0000VZ\u0003\u001d\u000e\u0000WZ\u0003\u001f\u000f"+
		"\u0000XZ\u0003!\u0010\u0000YV\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000"+
		"\u0000YX\u0001\u0000\u0000\u0000Z\u001c\u0001\u0000\u0000\u0000[]\u0005"+
		"-\u0000\u0000\\[\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000]_"+
		"\u0001\u0000\u0000\u0000^`\u0007\u0003\u0000\u0000_^\u0001\u0000\u0000"+
		"\u0000`a\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000ab\u0001\u0000"+
		"\u0000\u0000b\u001e\u0001\u0000\u0000\u0000cd\u00050\u0000\u0000de\u0005"+
		"x\u0000\u0000eg\u0001\u0000\u0000\u0000fh\u0007\u0004\u0000\u0000gf\u0001"+
		"\u0000\u0000\u0000hi\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000"+
		"ij\u0001\u0000\u0000\u0000j \u0001\u0000\u0000\u0000kl\u00050\u0000\u0000"+
		"lm\u0005b\u0000\u0000mo\u0001\u0000\u0000\u0000np\u0007\u0005\u0000\u0000"+
		"on\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000"+
		"\u0000qr\u0001\u0000\u0000\u0000r\"\u0001\u0000\u0000\u0000su\u0007\u0003"+
		"\u0000\u0000ts\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vt\u0001"+
		"\u0000\u0000\u0000vw\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000"+
		"x|\u0005.\u0000\u0000y{\u0007\u0003\u0000\u0000zy\u0001\u0000\u0000\u0000"+
		"{~\u0001\u0000\u0000\u0000|z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000"+
		"\u0000}$\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000\u007f\u0083"+
		"\u0005\"\u0000\u0000\u0080\u0082\u0003\'\u0013\u0000\u0081\u0080\u0001"+
		"\u0000\u0000\u0000\u0082\u0085\u0001\u0000\u0000\u0000\u0083\u0081\u0001"+
		"\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0086\u0001"+
		"\u0000\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0086\u0087\u0005"+
		"\"\u0000\u0000\u0087&\u0001\u0000\u0000\u0000\u0088\u008b\u0007\u0006"+
		"\u0000\u0000\u0089\u008b\u0003)\u0014\u0000\u008a\u0088\u0001\u0000\u0000"+
		"\u0000\u008a\u0089\u0001\u0000\u0000\u0000\u008b(\u0001\u0000\u0000\u0000"+
		"\u008c\u008d\u0005\\\u0000\u0000\u008d\u008e\u0007\u0007\u0000\u0000\u008e"+
		"*\u0001\u0000\u0000\u0000\u008f\u0092\u0003-\u0016\u0000\u0090\u0092\u0003"+
		"/\u0017\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0090\u0001\u0000"+
		"\u0000\u0000\u0092,\u0001\u0000\u0000\u0000\u0093\u0094\u0005t\u0000\u0000"+
		"\u0094\u0095\u0005r\u0000\u0000\u0095\u0096\u0005u\u0000\u0000\u0096\u0097"+
		"\u0005e\u0000\u0000\u0097.\u0001\u0000\u0000\u0000\u0098\u0099\u0005f"+
		"\u0000\u0000\u0099\u009a\u0005a\u0000\u0000\u009a\u009b\u0005l\u0000\u0000"+
		"\u009b\u009c\u0005s\u0000\u0000\u009c\u009d\u0005e\u0000\u0000\u009d0"+
		"\u0001\u0000\u0000\u0000\r\u0000KSY\\aiqv|\u0083\u008a\u0091\u0001\u0006"+
		"\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}